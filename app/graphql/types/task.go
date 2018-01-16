package types

import "github.com/graphql-go/graphql"

// Task type
var Task = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.Int},
		"userId":      &graphql.Field{Type: graphql.Int},
		"isComplete":  &graphql.Field{Type: graphql.Boolean},
		"title":       &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
		"deletedAt":   &graphql.Field{Type: graphql.DateTime},
	},
})
