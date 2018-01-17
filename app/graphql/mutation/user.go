package mutation

import (
	"errors"

	"github.com/charliekenney23/go-graphql-todo/app/auth"
	"github.com/charliekenney23/go-graphql-todo/app/graphql/types"
	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/graphql-go/graphql"
)

var updateUser = &graphql.Field{
	Type:        types.User,
	Description: "Update a task",
	Args: graphql.FieldConfigArgument{
		"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"email":     &graphql.ArgumentConfig{Type: graphql.String},
		"firstname": &graphql.ArgumentConfig{Type: graphql.String},
		"lastname":  &graphql.ArgumentConfig{Type: graphql.String},
		"username":  &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := auth.GetUserContextFrom(params.Context)
		uID := params.Args["id"].(int)

		user := &model.User{ID: uint(uID)}
		if err := shared.SharedApp.DB.Preload("Tasks").Find(&user).Error; err != nil {
			return nil, errors.New("User not found")
		}

		if !ctx.IsAdmin() && user.ID != ctx.ID {
			return nil, errors.New("Unauthorized to update user")
		}

		email, ok := params.Args["email"].(string)
		if ok {
			user.Email = email
		}

		firstname, ok := params.Args["firstname"].(string)
		if ok {
			user.Firstname = firstname
		}

		lastname, ok := params.Args["lastname"].(string)
		if ok {
			user.Firstname = lastname
		}

		username, ok := params.Args["username"].(string)
		if ok {
			user.Username = username
		}

		shared.SharedApp.DB.Update(&user)

		return user, nil
	},
}
