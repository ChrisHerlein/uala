package injectors

import (
	"github.com/ChrisHerlein/uala/tuala-api-users/src/handlers"
)

type Handlers struct {
	Users handlers.Users
}

func GetHandlers(services *Services) *Handlers {
	return &Handlers{
		handlers.NewUsers(services.Users),
	}
}
