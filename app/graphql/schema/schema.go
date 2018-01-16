package schema

import (
	"github.com/charliekenney23/go-graphql-todo/app/graphql/mutation"
	"github.com/charliekenney23/go-graphql-todo/app/graphql/query"

	"github.com/graphql-go/graphql"
)

// Schema is root GraphQL schema representation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query.Root,
	Mutation: mutation.Root,
})
