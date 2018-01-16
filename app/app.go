package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// side effect import for MYSQL driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/charliekenney23/go-graphql-todo/app/db"
	"github.com/charliekenney23/go-graphql-todo/app/handler"
	"github.com/charliekenney23/go-graphql-todo/app/middleware"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/charliekenney23/go-graphql-todo/config"
)

// SharedApp is the shared instance of the
// main app
var SharedApp *App

// App provides interface to extend the shared
// App type with the methods necessary for
// bootstrapping
type App shared.App

func init() {
	conf := config.Shared

	router := gin.Default()

	database, err := gorm.Open(conf.DB.Dialect, conf.DB.URI())
	if err != nil {
		log.Fatalf("Error opening DB: %v\n", err)
	}
	database = db.Migrate(database)

	SharedApp = newApp(shared.Initialize(router, database, conf))
	SharedApp.mountRoutes()
}

// Run runs the app engine and serves routes
func (a *App) Run(host string) {
	log.Fatal(a.Router.Run(host))
}

func (a *App) mountRoutes() {
	a.Router.POST("/user", handler.Register)
	a.Router.POST("/auth", handler.Authenticate)
	a.Router.GET("/graphql", middleware.RequireAuth, handler.GraphQL)
}

func newApp(a *shared.App) *App {
	return &App{a.Router, a.DB, a.Config}
}
