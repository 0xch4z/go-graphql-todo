package types

import "github.com/graphql-go/graphql"

// User type
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"email":     &graphql.Field{Type: graphql.String},
		"firstname": &graphql.Field{Type: graphql.String},
		"lastname":  &graphql.Field{Type: graphql.String},
		"username":  &graphql.Field{Type: graphql.String},
		"tasks":     &graphql.Field{Type: graphql.NewList(Task)},
	},
})
