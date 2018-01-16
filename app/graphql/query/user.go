package query

import (
	"errors"
	"fmt"

	"github.com/charliekenney23/go-graphql-complex/app/graphql/types"
	"github.com/charliekenney23/go-graphql-complex/app/model"
	"github.com/charliekenney23/go-graphql-complex/app/shared"
	"github.com/graphql-go/graphql"
)

var getUser = &graphql.Field{
	Type:        types.User,
	Description: "Get user by username",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var user model.User

		username := params.Args["username"].(string)

		if err := shared.SharedApp.DB.Preload("Tasks").Find(&user, "username = ?", username).Error; err != nil {
			return nil, errors.New("Could not find user")
		}

		fmt.Println("id => ", user.ID)

		return user, nil
	},
}
