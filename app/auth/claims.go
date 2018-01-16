package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims are the payload claimed from a
// valid JWT token
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	ID       uint   `json:"id"`
	Role     string `json:"role"`
}
