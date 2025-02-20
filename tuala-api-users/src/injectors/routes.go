package injectors

import (
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/routes"
)

func SetRoutes(router *fiber.App, handlers *Handlers, middlewares *Middlewares) {
	users := routes.NewUsers()
	users.Add(router, handlers.Users, middlewares.Auth)
}
