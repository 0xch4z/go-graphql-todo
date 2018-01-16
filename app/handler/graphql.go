package handler

import "github.com/gin-gonic/gin"

func GraphQL(c *gin.Context) {
	c.JSON(200, &gin.H{
		"protected": true,
	})
}
