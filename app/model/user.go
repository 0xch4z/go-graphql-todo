package model

import (
	"github.com/charliekenney23/go-graphql-complex/app/shared"
)

// User type
type User struct {
	Model
	Email     string `json:"email" gorm:"unique"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" gorm:"unique"`
	Password  []byte `json:"password"`
	Role      string `json:"role" gorm:"type:ENUM('user','admin');default:'user'"`
	Tasks     []Task `json:"tasks"`
}

// FindUserByUsername finds a given user by username or
// rethrows an error
func FindUserByUsername(username string) (*User, error) {
	var user *User

	if err := shared.SharedApp.DB.Where(&User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
