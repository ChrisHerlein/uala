package services

import (
	"context"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/repositories"
)

type Feed interface {
	Recent(ctx context.Context, page int) ([]models.Content, error)
}

type feed struct {
	cache repositories.Cache
}

func (f *feed) Recent(ctx context.Context, page int) ([]models.Content, error) {
	userName := ctx.Value(enums.CtxUserName).(string)
	return f.cache.GetFeed(userName, page)
}

func NewFeed(cache repositories.Cache) *feed {
	return &feed{cache}
}
