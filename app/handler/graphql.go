package handler

import (
	"log"
	"net/http"

	"github.com/charliekenney23/go-graphql-complex/app/graphql/schema"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// GraphQL executes a graphql query
func GraphQL(c *gin.Context) {
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: c.Query("query"),
		Context:       c.Request.Context(),
	})

	if len(result.Errors) > 0 {
		log.Printf("Unexpected error(s): %v\n", result.Errors)
	}

	c.JSON(http.StatusOK, result)
}
