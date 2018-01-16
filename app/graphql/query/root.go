package query

import "github.com/graphql-go/graphql"

// Root Query
var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getTask":     getTask,
		"getAllTasks": getAllTasks,
		"getUser":     getUser,
	},
})
