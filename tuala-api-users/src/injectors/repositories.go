package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-users/src/repositories"
)

type Repositories struct {
	DB    repositories.DB
	Cache repositories.Cache
}

func GetRepositories(conns *Connections) *Repositories {
	return &Repositories{
		DB:    repositories.NewDB(conns.PostgreSQL),
		Cache: repositories.NewWorkerCache(conns.Beanstalk),
	}
}
