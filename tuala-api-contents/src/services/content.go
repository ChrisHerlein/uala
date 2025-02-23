package services

import (
	"context"
	"fmt"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/repositories"
)

type Content interface {
	Read(userName string, page int) ([]models.Content, error)
	Create(ctx context.Context, text string) (*models.Content, error)
}

type content struct {
	db    repositories.DB
	cache repositories.Cache
}

func (c *content) Create(ctx context.Context, text string) (*models.Content, error) {
	if len(text) > 240 {
		return nil, fmt.Errorf("%w text is longer than 240", enums.Err400)
	}
	tweet := &models.Content{
		AuthorName: ctx.Value(enums.CtxUserName).(string),
		Text:       text,
	}

	if err := c.db.CreateTweet(tweet); err != nil {
		return nil, err
	}

	go c.cache.RecreateFeed(ctx.Value(enums.CtxUserID).(uint), tweet.ID)

	return tweet, nil
}

func (c *content) Read(userName string, page int) ([]models.Content, error) {
	return c.db.ReadTweets(userName, page)
}

func NewContent(db repositories.DB, cache repositories.Cache) *content {
	return &content{db, cache}
}
