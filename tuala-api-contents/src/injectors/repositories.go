package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-contents/src/repositories"
	usersRepository "github.com/ChrisHerlein/uala/tuala-api-users/src/repositories"
)

type Repositories struct {
	UsersDB usersRepository.DB
	DB      repositories.DB
	Cache   repositories.Cache
}

func GetRepositories(conns *Connections) *Repositories {
	return &Repositories{
		UsersDB: usersRepository.NewDB(conns.PostgreSQL),
		DB:      repositories.NewDB(conns.PostgreSQL),
		Cache:   repositories.NewCache(conns.Beanstalk, conns.Redis),
	}
}
