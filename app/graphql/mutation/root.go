package mutation

import "github.com/graphql-go/graphql"

// Root Mutation
var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTask": createTask,
	},
})
