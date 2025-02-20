package injectors

import (
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/middlewares"
)

type Middlewares struct {
	Auth func(c fiber.Ctx) error
}

func GetMiddlewares(repositories *Repositories) *Middlewares {
	return &Middlewares{
		Auth: middlewares.Auth(repositories.DB),
	}
}
