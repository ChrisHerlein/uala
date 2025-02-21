package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/handlers"
)

type feedRoutes struct{}

func (cr *feedRoutes) Add(router *fiber.App, handler handlers.Feed, authMid func(c fiber.Ctx) error) {
	path := router.Group(enums.FeedRoutes)
	path.Get(enums.FeedRecentRoute, handler.Recent(), authMid)
}

func NewFeed() *feedRoutes {
	return &feedRoutes{}
}

type contentRoutes struct{}

func (cr *contentRoutes) Add(router *fiber.App, handler handlers.Content, authMid func(c fiber.Ctx) error) {
	path := router.Group(enums.ContentRoutes)
	path.Get(enums.ContentGet, handler.Get())
	path.Post(enums.ContentCreate, handler.Create(), authMid)
}

func NewContent() *contentRoutes {
	return &contentRoutes{}
}
