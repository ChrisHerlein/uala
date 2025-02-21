package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/handlers"
)

type Handlers struct {
	Feed    handlers.Feed
	Content handlers.Content
}

func GetHandlers(services *Services) *Handlers {
	return &Handlers{
		Content: handlers.NewContent(services.Content),
		Feed:    handlers.NewFeed(services.Feed),
	}
}
