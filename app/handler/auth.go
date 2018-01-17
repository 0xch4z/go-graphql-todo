package handler

import (
	"encoding/json"
	"net/http"

	"github.com/charliekenney23/go-graphql-todo/app/auth"
	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/gin-gonic/gin"
)

// AuthBuffer type represents object of values
// necessary for authentication
type AuthBuffer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticate a given user and return a JWT token
func Authenticate(c *gin.Context) {
	ab := &AuthBuffer{}

	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(ab); err != nil {
		abortWithInternalServerError(c, err)
		return
	}

	user := &model.User{}

	if err := shared.SharedApp.DB.Find(&user, "username = ?", ab.Username).Error; err != nil {
		abortWithNotFoundError(c, "User")
		return
	}

	match := auth.CheckPassword(user.Password, ab.Password)
	if !match {
		abortWithUnauthorizedError(c, "incorrect username or password")
		return
	}

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
