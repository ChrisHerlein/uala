package services

import (
	"context"

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
