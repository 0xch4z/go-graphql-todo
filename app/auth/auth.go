package auth

import (
	"time"

	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// GenerateToken generates a new JWT token with
// claims for a given user
func GenerateToken(user *model.User) (*string, error) {
	tok := jwt.New(jwt.SigningMethodRS256)

	now := time.Now()

	tok.Claims = Claims{
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * time.Duration(24*90)).Unix(),
			IssuedAt:  now.Unix(),
		},
		user.Username,
		user.ID,
		user.Role,
	}

	tokString, err := tok.SignedString(shared.SharedApp.Config.Crypto.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &tokString, nil
}

// HashPassword generates a hash for a given password
func HashPassword(password string) (*[]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &hash, nil
}

// CheckPassword checks to see if a password hash matches
// an attempt
func CheckPassword(password []byte, attempt string) bool {
	err := bcrypt.CompareHashAndPassword(password, []byte(attempt))
	return err == nil
}
