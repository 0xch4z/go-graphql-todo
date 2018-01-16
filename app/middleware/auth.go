package middleware

import (
	"net/http"

	"github.com/charliekenney23/go-graphql-complex/app/auth"
	"github.com/charliekenney23/go-graphql-complex/app/shared"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RequireAuth requires authorization and pushes the
// claims to the context. If the JWT is not valid or
// is not present, a 401 will be thrown
func RequireAuth(c *gin.Context) {
	tok := c.Request.Header.Get("Authorization")

	token, err := jwt.ParseWithClaims(tok, &auth.Claims{}, func(tok *jwt.Token) (interface{}, error) {
		return shared.SharedApp.Config.Crypto.PublicKey, nil
	})
	if err != nil {
		abortUnauthorized(c)
		return
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		c.Set("id", int(claims.ID))
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
	} else {
		abortUnauthorized(c)
	}
}

func abortUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, &gin.H{
		"error":   true,
		"message": "Unauthorized",
	})
}
