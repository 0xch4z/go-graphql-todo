package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func abortWithError(c *gin.Context, code int, msg ...string) {
	var m string

	if len(msg) == 0 {
		m = "Something Broke ):"
	} else {
		m = strings.Join(msg, " ")
	}

	c.AbortWithStatusJSON(code, &gin.H{
		"error":   true,
		"message": m,
	})
}

func abortWithNotFoundError(c *gin.Context, resource ...string) {
	var m string

	if len(resource) == 0 {
		m = "Not found"
	} else {
		m = fmt.Sprintf("%s not found", strings.Join(resource, ""))
	}

	abortWithError(c, http.StatusNotFound, m)
}

func abortWithInternalServerError(c *gin.Context, err error) {
	fmt.Printf("Error: %v", err)
	abortWithError(c, http.StatusInternalServerError)
}

func abortWithUnauthorizedError(c *gin.Context, reason string) {
	abortWithError(c, http.StatusUnauthorized, reason)
}
