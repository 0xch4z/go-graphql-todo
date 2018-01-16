package mutation

import (
	"errors"

	"github.com/charliekenney23/go-graphql-complex/app/graphql/types"
	"github.com/charliekenney23/go-graphql-complex/app/model"
	"github.com/charliekenney23/go-graphql-complex/app/shared"
	"github.com/graphql-go/graphql"
)

var createTask = &graphql.Field{
	Type:        types.Task,
	Description: "Create a new task",
	Args: graphql.FieldConfigArgument{
		"title":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"description": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Context.Value("id").(int)
		title, _ := params.Args["title"].(string)
		desc, _ := params.Args["description"].(string)

		newTask := model.Task{Title: title, Description: desc, UserID: uint(id)}

		tx := shared.SharedApp.DB.Begin()
		if err := tx.Create(&newTask).Error; err != nil {
			return nil, errors.New("Could not create task")
		}
		tx.Commit()

		return newTask, nil
	},
}
