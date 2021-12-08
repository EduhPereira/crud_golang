package migrations

import (
	"github.com/EduhPereira/crud_golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}