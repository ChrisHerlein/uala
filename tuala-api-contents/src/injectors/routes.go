package injectors

import (
	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/routes"
)

func SetRoutes(router *fiber.App, handlers *Handlers, middlewares *Middlewares) {
	feed := routes.NewFeed()
	feed.Add(router, handlers.Feed, middlewares.Auth)

	content := routes.NewContent()
	content.Add(router, handlers.Content, middlewares.Auth)
}
