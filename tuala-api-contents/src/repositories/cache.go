package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	beanstalk "github.com/beanstalkd/go-beanstalk"
	"github.com/redis/go-redis/v9"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

type message struct {
	UserID   uint `json:"userId"`
	PageRead int  `json:"pageRead"`
}

type WorkerCache interface {
	RecreateFeed(userID uint) error
	MarkPageRead(userID uint, pageNumber int)
}

type ReadCache interface {
	GetFeed(userID uint) ([]models.FeedPage, error)
}

type Cache interface {
	WorkerCache
	ReadCache
}

type cache struct {
	WorkerCache
	ReadCache
}

type workerCache struct {
	beanstalk  *beanstalk.Conn
	tubeFollow *beanstalk.Tube
	tubeRead   *beanstalk.Tube
}

func (wc *workerCache) RecreateFeed(userID uint) error {
	msg := message{UserID: userID}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = wc.tubeFollow.Put(msgBytes, 1, 0, 0)
	return err
}

func (wc *workerCache) MarkPageRead(userID uint, pageNumber int) {
	msg := message{UserID: userID, PageRead: pageNumber}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return
	}
	wc.tubeRead.Put(msgBytes, 1, 0, 0)
	return
}

func NewWorkerCache(bc *beanstalk.Conn) *workerCache {
	return &workerCache{
		beanstalk:  bc,
		tubeFollow: beanstalk.NewTube(bc, enums.QueueRecreateFeedNewContent),
		tubeRead:   beanstalk.NewTube(bc, enums.QueueRecreateFeedPageRead),
	}
}

type redisCache struct {
	rc *redis.Client
}

func (rc *redisCache) GetFeed(userID uint) ([]models.FeedPage, error) {
	fc, err := rc.rc.Get(context.TODO(), fmt.Sprintf("%d-control", userID)).Bytes()
	if err != nil {
		return nil, err
	}

	var ctrlDoc models.Control
	err = json.Unmarshal(fc, &ctrlDoc)
	if err != nil {
		return nil, err
	}

	var pages = make([]models.FeedPage, 0)
	for i := 0; i < 2 && i < ctrlDoc.MostRecent; i++ {
		var page models.FeedPage
		fp, err := rc.rc.Get(context.TODO(), fmt.Sprintf("%d-%d", userID, ctrlDoc.MostRecent-i)).Bytes()
		if err != nil {
			continue
		}
		err = json.Unmarshal(fp, &page)
		if err != nil {
			continue
		}
		pages = append(pages, page)
	}

	return pages, nil
}

func NewRedisCache(rc *redis.Client) *redisCache {
	return &redisCache{rc}
}

func NewCache(bc *beanstalk.Conn, rc *redis.Client) *cache {
	return &cache{
		NewWorkerCache(bc),
		NewRedisCache(rc),
	}
}
