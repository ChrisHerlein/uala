package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/services"
)

type Services struct {
	Feed    services.Feed
	Content services.Content
}

func GetServices(repos *Repositories) *Services {
	return &Services{
		Feed:    services.NewFeed(repos.Cache),
		Content: services.NewContent(repos.DB, repos.Cache),
	}
}
