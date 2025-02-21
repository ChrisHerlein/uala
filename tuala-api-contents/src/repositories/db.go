package repositories

import (
	"time"

	"gorm.io/gorm"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

const pageLength = 5

type DB interface {
	CreateTweet(tweet *models.Content) error
	ReadTweets(userName string, page int) ([]models.Content, error)
}

type pgdb struct {
	db *gorm.DB
}

func (pg *pgdb) CreateTweet(tweet *models.Content) error {
	tweet.CreatedAt = time.Now()
	tweet.UpdatedAt = time.Now()

	result := pg.db.Create(tweet)

	return result.Error
}

func (pg *pgdb) ReadTweets(userName string, page int) ([]models.Content, error) {
	var content []models.Content

	res := pg.db.Where("author_name = ?", userName).
		Order("created_at desc").
		Limit(pageLength).
		Offset(pageLength * page).
		Find(&content)

	return content, res.Error
}

func (pg *pgdb) Filter(query string, args []interface{}) ([]models.Content, error) {
	var content []models.Content

	res := pg.db.
		Where(query, args...).
		Find(&content)

	return content, res.Error
}

func NewDB(db *gorm.DB) *pgdb {
	return &pgdb{
		db: db,
	}
}
