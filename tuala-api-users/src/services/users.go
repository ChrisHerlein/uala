package services

import (
	"context"
	"fmt"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/models"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/repositories"
)

type Users interface {
	Get(name string) (*models.User, error)
	Create(name, password string) (*models.User, error)
	Follow(ctx context.Context, to string) (*models.User, error)
	Unfollow(ctx context.Context, to string) (*models.User, error)
}

type users struct {
	db    repositories.DB
	cache repositories.Cache
}

func (u *users) Create(name, password string) (*models.User, error) {
	if _, err := u.db.Get(name, ""); err == nil { // means user name already exists
		return nil, fmt.Errorf("%w user alred exists", enums.Err409)
	}

	user := &models.User{
		Name:     name,
		Password: password,
	}

	return user, u.db.CreateUser(user)
}

func (u *users) Get(name string) (*models.User, error) {
	return u.db.Get(name, "")
}

func (u *users) Follow(ctx context.Context, to string) (*models.User, error) {
	toFollow, err := u.db.Get(to, "")
	if err != nil {
		return nil, fmt.Errorf("%w user to follow does not exists", enums.Err404)
	}

	fromID := ctx.Value(enums.CtxUserID).(uint)
	err = u.db.Follow(fromID, toFollow.ID)
	if err != nil {
		return nil, fmt.Errorf("%w %w", enums.Err503, err)
	}

	err = u.cache.RecreateFeed(fromID)
	if err != nil {
		return nil, fmt.Errorf("%w %w", enums.Err503, err)
	}

	fromName := ctx.Value(enums.CtxUserName).(string)
	return u.db.Get(fromName, "")
}

func (u *users) Unfollow(ctx context.Context, to string) (*models.User, error) {
	toUnfollow, err := u.db.Get(to, "")
	if err != nil {
		return nil, fmt.Errorf("%w user to unfollow does not exists", enums.Err404)
	}

	fromID := ctx.Value(enums.CtxUserID).(uint)
	err = u.db.Unfollow(fromID, toUnfollow.ID)
	if err != nil {
		return nil, fmt.Errorf("%w %w", enums.Err503, err)
	}

	fromName := ctx.Value(enums.CtxUserName).(string)
	return u.db.Get(fromName, "")
}

func NewUsers(db repositories.DB, cache repositories.Cache) *users {
	return &users{db, cache}
}
