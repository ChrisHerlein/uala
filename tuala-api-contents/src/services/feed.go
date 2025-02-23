package services

import (
	"context"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/repositories"
)

type Feed interface {
	Recent(ctx context.Context) ([]models.Content, error)
}

type feed struct {
	cache repositories.Cache
}

func (f *feed) Recent(ctx context.Context) ([]models.Content, error) {
	userID := ctx.Value(enums.CtxUserID).(uint)
	pages, err := f.cache.GetFeed(userID)
	if err != nil {
		// to avoid empty feed causes an error
		return []models.Content{}, nil
	}

	var content = make([]models.Content, 0)
	for i := 0; i < len(pages); i++ {
		content = append(content, pages[i].Content...)
		f.cache.MarkPageRead(userID, pages[i].Order)
	}

	return content, nil
}

func NewFeed(cache repositories.Cache) *feed {
	return &feed{cache}
}
