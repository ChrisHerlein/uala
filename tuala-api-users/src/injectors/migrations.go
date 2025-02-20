package injectors

import (
	"gorm.io/gorm"

	"github.com/ChrisHerlein/uala/tuala-api-users/src/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Follow{})
}
