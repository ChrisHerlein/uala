package main

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	contentModels "github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
	contentRepo "github.com/ChrisHerlein/uala/tuala-api-contents/src/repositories"
	usersRepo "github.com/ChrisHerlein/uala/tuala-api-users/src/repositories"
)

const pageSize = 10

type engine struct {
	dbUsers    usersRepo.DB
	dbContents contentRepo.DB
	rh         *cacheHandler
	locks      sync.Map
}

func (e *engine) awaitAndProcess(messages chan message, process func(message)) {
	for {
		msg := <-messages
		go process(msg)
	}
}

func (e *engine) lockUser(userID uint) *sync.Mutex {
	mutex, ok := e.locks.Load(userID)
	if !ok {
		mutex = &sync.Mutex{}
		e.locks.Store(userID, mutex)
	}

	m := mutex.(*sync.Mutex)
	m.Lock()
	return m
}

func (e *engine) recreateByFollow(msg message) {
	m := e.lockUser(msg.UserID)
	defer m.Unlock()

	user, err := e.dbUsers.GetByID(msg.UserID)
	if err != nil {
		log.Default().Println("[engine][recreateByFollow][GetByID] Error:", err)
		return
	}

	// Get Last Followed
	lastFollowed := user.Follows[len(user.Follows)-1]

	userFollowed, err := e.dbUsers.GetByID(lastFollowed.Follows)
	if err != nil {
		log.Default().Println("[engine][recreateByFollow][GetByID] Error:", err)
		return
	}

	// Get last page of last followed
	lastContent, err := e.dbContents.ReadTweets(userFollowed.Name, 0)
	if err != nil {
		log.Default().Println("[engine][recreateByFollow][ReadTweets] Error:", err)
		return
	}

	// Get control doc
	ctrlDoc := e.getControlDoc(msg.UserID)

	// if last page not completed, get lastPage
	lastPage := &feedPage{}
	if ctrlDoc.SizeOfLast != pageSize && ctrlDoc.MostRecent != 0 {
		lastPage = e.getLastPage(msg.UserID, ctrlDoc.MostRecent)
		remainingLastPage := pageSize - len(lastPage.Content)
		startFrom := len(lastContent) - 1 - remainingLastPage
		if startFrom < 0 {
			startFrom = 0
		}
		remainingContent := lastContent[startFrom:len(lastContent)]
		content := make([]contentModels.Content, 0)
		content = append(content, remainingContent...)
		content = append(content, lastPage.Content...)
		lastPage.Content = content

		lastContent = lastContent[:startFrom]

		// update last page
		e.upsertPage(msg.UserID, lastPage)
		ctrlDoc.SizeOfLast = len(lastPage.Content)
	}

	// check if new page is needed
	if len(lastContent) > 0 {
		newPage := feedPage{
			UserName: lastPage.UserName,
			Order:    lastPage.Order + 1,
			Content:  make([]contentModels.Content, 0),
		}

		// fullfill new page
		// already sorted by newest
		for i := 0; i < len(lastContent); i++ {
			newPage.Content = append(newPage.Content, lastContent[i])
		}

		// update last page
		e.upsertPage(msg.UserID, &newPage)
		ctrlDoc.SizeOfLast = len(newPage.Content)
		ctrlDoc.MostRecent = newPage.Order
	}

	// update control doc
	e.upsertCtrlDoc(msg.UserID, ctrlDoc)
}

func (e *engine) getControlDoc(userID uint) *control {
	// if no present, return empty
	// that will be created
	ctrlDoc := control{}
	err := e.rh.get(fmt.Sprintf("%d-control", userID), &ctrlDoc)
	if err != nil {
		log.Default().Println("[engine][getControlDoc][get] Error:", err)
	}
	return &ctrlDoc
}

func (e *engine) getLastPage(userID uint, pageN int) *feedPage {
	page := feedPage{}
	err := e.rh.get(fmt.Sprintf("%d-%d", userID, pageN), &page)
	if err != nil {
		log.Default().Println("[engine][getLastPage][get] Error:", err)
	}
	return &page
}

func (e *engine) upsertPage(userID uint, page *feedPage) {
	err := e.rh.set(fmt.Sprintf("%d-%d", userID, page.Order), page)
	if err != nil {
		log.Default().Println("[engine][upsertPage][set] Error:", err)
	}
}

func (e *engine) upsertCtrlDoc(userID uint, ctrl *control) {
	err := e.rh.set(fmt.Sprintf("%d-control", userID), ctrl)
	if err != nil {
		log.Default().Println("[engine][upsertCtrlDoc][set] Error:", err)
	}
}

func (e *engine) recreateByContent(msg message) {
	// get users that follows author
	follows, err := e.dbUsers.GetFollowers(msg.UserID)
	if err != nil {
		log.Default().Println("[engine][recreateByContent][GetFollowers] Error:", err)
		return
	}

	// get new content
	contents, err := e.dbContents.Filter("id = ?", []interface{}{msg.ContentID})
	if err != nil {
		log.Default().Println("[engine][recreateByContent][contentRepo.Filter] Error:", err)
		return
	}

	// apply cache update
	for i := 0; i < len(follows); i++ {
		m := e.lockUser(follows[i].UserRefer)
		e.recreateUserFeedByContent(follows[i].UserRefer, contents[0])
		m.Unlock()
	}
}

func (e *engine) recreateUserFeedByContent(userID uint, content contentModels.Content) {
	// Get control doc
	ctrlDoc := e.getControlDoc(userID)

	// lets guess we need to create a new page
	var page = &feedPage{
		UserName: ctrlDoc.UserName,
		Order:    ctrlDoc.MostRecent + 1,
		Content:  make([]contentModels.Content, 0),
	}

	// if last page is not full and it exists
	if ctrlDoc.SizeOfLast < pageSize && ctrlDoc.MostRecent != 0 {
		page = e.getLastPage(userID, ctrlDoc.MostRecent)
	}

	// page.Content = append(page.Content, content)
	page.Content = append([]contentModels.Content{content}, page.Content...)
	e.upsertPage(userID, page)

	ctrlDoc.SizeOfLast = len(page.Content)
	ctrlDoc.MostRecent = page.Order
	e.upsertCtrlDoc(userID, ctrlDoc)
}

func (e *engine) pageRead(msg message) {
	m := e.lockUser(msg.UserID)
	defer m.Unlock()

	// Get control doc
	ctrlDoc := e.getControlDoc(msg.UserID)

	// remove read page
	err := e.rh.remove(fmt.Sprintf("%d-%d", msg.UserID, msg.PageRead))
	if err != nil {
		log.Default().Println("[engine][pageRead][remove] Error:", err)
		return
	}

	// update control doc
	ctrlDoc.MostRecent = msg.PageRead - 1
	e.upsertCtrlDoc(msg.UserID, ctrlDoc)
}

func newEngine(cfg *Config, rh *cacheHandler) (*engine, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDb, cfg.PgPort,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &engine{
		dbContents: contentRepo.NewDB(db),
		dbUsers:    usersRepo.NewDB(db),
		rh:         rh,
	}, nil
}
