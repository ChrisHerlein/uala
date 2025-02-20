package repositories

import (
	"encoding/json"

	beanstalk "github.com/beanstalkd/go-beanstalk"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
)

type message struct {
	UserID uint `json:"userId"`
}

type Cache interface {
	RecreateFeed(userId uint) error
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
		tubeFollow: beanstalk.NewTube(bc, enums.QueueRecreateFeedFollow),
	}
}
