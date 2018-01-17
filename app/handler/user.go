package handler

import (
	"encoding/json"
	"net/http"

	"github.com/charliekenney23/go-graphql-todo/app/auth"
	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/gin-gonic/gin"
)

// UserBuffer type represents a new user graph
// for registration
type UserBuffer struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// Register registers a new user
func Register(c *gin.Context) {
	ub := &UserBuffer{}

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(ub); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"error": true,
		})
		return
	}
	defer c.Request.Body.Close()

	hash, err := auth.HashPassword(ub.Password)
	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	user := &model.User{
		Email:     ub.Email,
		Firstname: ub.Firstname,
		Lastname:  ub.Lastname,
		Username:  ub.Username,
		Password:  *hash,
		Role:      "user",
	}

	tx := shared.SharedApp.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		abortWithInternalServerError(c, err)
		return
	}
	tx.Commit()

	tok, err := auth.GenerateToken(user)
	if err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &gin.H{
		"success": true,
		"token":   tok,
	})
}
