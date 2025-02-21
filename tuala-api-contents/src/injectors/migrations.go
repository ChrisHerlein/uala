package injectors

import (
	"gorm.io/gorm"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Content{})
}
