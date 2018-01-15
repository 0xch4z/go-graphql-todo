package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// side effect import for MYSQL driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/charliekenney23/go-graphql-complex/app/shared"
	"github.com/charliekenney23/go-graphql-complex/config"
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
	db, err := gorm.Open(conf.DB.Dialect, conf.DB.URI())
	if err != nil {
		log.Fatalf("Error opening DB: %v\n", err)
	}

	SharedApp = newApp(shared.Initialize(router, db, conf))
	SharedApp.mountRoutes()
}

// Run runs the app engine and serves routes
func (a *App) Run(host string) {
	log.Fatal(a.Router.Run(host))
}

func (a *App) mountRoutes() {
	a.Router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, &map[string]interface{}{
			"success": true,
		})
	})
}

func newApp(a *shared.App) *App {
	return &App{a.Router, a.DB, a.Config}
}
