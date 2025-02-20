package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/handlers"
)

type userRoutes struct{}

func (ur *userRoutes) Add(router *fiber.App, handler handlers.Users, authMid func(c fiber.Ctx) error) {
	path := router.Group(enums.UsersRoutes)
	path.Get(enums.UsersGet, handler.Get())
	path.Post(enums.UsersCreate, handler.Create())
	path.Post(enums.UsersFollow, handler.Follow(), authMid)
	path.Post(enums.UsersUnfollow, handler.Unfollow(), authMid)
}

func NewUsers() *userRoutes {
	return &userRoutes{}
}
