package db

import (
	"github.com/charliekenney23/go-graphql-complex/app/model"
	"github.com/jinzhu/gorm"
)

// Migrate migrates the DB if needed
func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{})
	return db
}
