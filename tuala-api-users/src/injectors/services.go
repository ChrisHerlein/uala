package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-users/src/services"
)

type Services struct {
	Users services.Users
}

func GetServices(repos *Repositories) *Services {
	return &Services{
		Users: services.NewUsers(repos.DB, repos.Cache),
	}
}
