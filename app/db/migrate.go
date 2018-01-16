package db

import (
	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/jinzhu/gorm"
)

// Migrate migrates the DB if needed
func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{}, &model.Task{})
	db.Model(&model.User{}).Related(&model.Task{})
	return db
}
