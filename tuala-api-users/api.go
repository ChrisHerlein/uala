package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/config"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/injectors"
)

func main() {
	/* Load Environment Variables */
	cfg := &config.Config{}
	cfg.Init()

	/* Instantiate server */
	app := fiber.New()

	/* Add repositories and services */
	connections := injectors.GetConnections(cfg)
	defer func() {
		db, err := connections.PostgreSQL.DB()
		if err != nil {
			return
		}
		db.Close()
	}()
	defer connections.Redis.Close()
	defer connections.Beanstalk.Close()

	injectors.RunMigrations(connections.PostgreSQL)
	repositories := injectors.GetRepositories(connections)
	services := injectors.GetServices(repositories)
	handlers := injectors.GetHandlers(services)
	middlewares := injectors.GetMiddlewares(repositories)

	injectors.SetRoutes(app, handlers, middlewares)

	err := app.Listen(fmt.Sprintf("0.0.0.0:%s", cfg.Port))
	if err != nil {
		panic(err)
	}
}
