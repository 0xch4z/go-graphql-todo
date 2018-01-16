package query

import "github.com/graphql-go/graphql"

// Root Query
var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getUser": getUser,
	},
})
