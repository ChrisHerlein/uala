package repositories

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/enums"
	"github.com/ChrisHerlein/uala/tuala-api-users/src/models"
)

var UniqueViolationCode = "23505"

type DB interface {
	CreateUser(user *models.User) error
	Get(name, password string) (*models.User, error)
	Follow(from, to uint) error
	Unfollow(from, to uint) error
}

type pgdb struct {
	db *gorm.DB
}

func (pg *pgdb) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result := pg.db.Create(user)

	return result.Error
}

func (pg *pgdb) Get(name, password string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Password: password,
	}

	res := pg.db.Where(user).Preload("Follows").First(user)
	return user, res.Error
}

func (pg *pgdb) Follow(from, to uint) error {
	uf := &models.Follow{
		UserRefer: from,
		Follows:   to,
	}
	res := pg.db.Create(uf)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), UniqueViolationCode) {
			return fmt.Errorf("%w user already follows", enums.Err409)
		}
	}
	return res.Error
}

func (pg *pgdb) Unfollow(from, to uint) error {
	uf := &models.Follow{}
	res := pg.db.Delete(uf, "user_refer = ? AND follows = ?", from, to)
	return res.Error
}

func NewDB(db *gorm.DB) *pgdb {
	return &pgdb{
		db: db,
	}
}
