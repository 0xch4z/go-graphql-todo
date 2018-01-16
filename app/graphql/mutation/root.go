package mutation

import "github.com/graphql-go/graphql"

// Root Mutation
var Root = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// task mutations
		"createTask":   createTask,
		"deleteTask":   deleteTask,
		"updateTask":   updateTask,
		"completeTask": completeTask,
		"undoTask":     undoTask,
		// user mutations
		"updateUser": updateUser,
	},
})
