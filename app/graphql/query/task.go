package query

import (
	"errors"

	"github.com/charliekenney23/go-graphql-complex/app/graphql/types"
	"github.com/charliekenney23/go-graphql-complex/app/model"
	"github.com/charliekenney23/go-graphql-complex/app/shared"
	"github.com/graphql-go/graphql"
)

var getTask = &graphql.Field{
	Type:        types.Task,
	Description: "Get task by ID",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// ctx := auth.GetUserContextFrom(params.Context)
		id := params.Args["id"].(int)

		task := model.Task{ID: uint(id)}
		if err := shared.SharedApp.DB.Find(&task).Error; err != nil {
			return nil, errors.New("Could not find task")
		}

		return task, nil
	},
}

var getAllTasks = &graphql.Field{
	Type:        graphql.NewList(types.Task),
	Description: "Get task by ID",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		tasks := []model.Task{}

		if err := shared.SharedApp.DB.Find(&tasks).Error; err != nil {
			return nil, errors.New("Could not resolve tasks")
		}

		return tasks, nil
	},
}
