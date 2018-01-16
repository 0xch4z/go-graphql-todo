package mutation

import (
	"errors"

	"github.com/charliekenney23/go-graphql-complex/app/auth"
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

var deleteTask = &graphql.Field{
	Type:        types.Task,
	Description: "Delete a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Context.Value("id").(int)
		isAdmin := params.Context.Value("role").(string) == "admin"
		tID := params.Context.Value("id").(int)

		task := &model.Task{ID: uint(tID)}
		if err := shared.SharedApp.DB.Find(&task).Error; err != nil {
			return nil, errors.New("Task not found")
		}

		if !isAdmin && task.UserID != uint(id) {
			return nil, errors.New("Unauthorized to delete task")
		}

		shared.SharedApp.DB.Delete(&task)
		return task, nil
	},
}

var updateTask = &graphql.Field{
	Type:        types.Task,
	Description: "Update a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		// "email":     &graphql.Field{Type: graphql.String},
		// "firstname": &graphql.Field{Type: graphql.String},
		// "lastname":  &graphql.Field{Type: graphql.String},
		// "username":  &graphql.Field{Type: graphql.String},
		"isComplete":  &graphql.ArgumentConfig{Type: graphql.Boolean},
		"title":       &graphql.ArgumentConfig{Type: graphql.String},
		"description": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := params.Context.Value("id").(int)
		isAdmin := params.Context.Value("role").(string) == "admin"
		tID := params.Args["id"].(int)

		task := &model.Task{ID: uint(tID)}
		if err := shared.SharedApp.DB.Find(&task).Error; err != nil {
			return nil, errors.New("Task not found")
		}

		if !isAdmin && task.UserID != uint(id) {
			return nil, errors.New("Unauthorized to update task")
		}

		isComplete, ok := params.Args["isComplete"].(bool)
		if ok {
			task.IsComplete = isComplete
		}

		title, ok := params.Args["title"].(string)
		if ok {
			task.Title = title
		}

		description, ok := params.Args["description"].(string)
		if ok {
			task.Description = description
		}

		shared.SharedApp.DB.Update(&task)

		return task, nil
	},
}

var completeTask = &graphql.Field{
	Type:        types.Task,
	Description: "Complete a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := auth.GetUserContextFrom(params.Context)
		tID := params.Args["id"].(int)

		task := &model.Task{ID: uint(tID)}
		if err := shared.SharedApp.DB.Find(&task).Error; err != nil {
			return nil, errors.New("Task not found")
		}

		if ctx.IsAdmin() && task.UserID != uint(ctx.ID) {
			return nil, errors.New("Unauthorized to complete task")
		}

		task.IsComplete = true

		shared.SharedApp.DB.Update(&task)

		return task, nil
	},
}

var undoTask = &graphql.Field{
	Type:        types.Task,
	Description: "Undo a task",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := auth.GetUserContextFrom(params.Context)
		tID := params.Args["id"].(int)

		task := &model.Task{ID: uint(tID)}
		if err := shared.SharedApp.DB.Find(&task).Error; err != nil {
			return nil, errors.New("Task not found")
		}

		if ctx.IsAdmin() && task.UserID != uint(ctx.ID) {
			return nil, errors.New("Unauthorized to undo task")
		}

		task.IsComplete = false

		shared.SharedApp.DB.Update(&task)

		return task, nil
	},
}
