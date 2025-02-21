package repositories

import (
	"encoding/json"

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
}

type ReadCache interface {
	GetFeed(userName string, page int) ([]models.Content, error)
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

func NewWorkerCache(bc *beanstalk.Conn) *workerCache {
	return &workerCache{
		beanstalk:  bc,
		tubeFollow: beanstalk.NewTube(bc, enums.QueueRecreateFeedNewContent),
	}
}

type redisCache struct {
	rc *redis.Client
}

func (rc *redisCache) GetFeed(userName string, page int) ([]models.Content, error) {
	return nil, nil
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
