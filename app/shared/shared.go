package shared

import (
	"github.com/charliekenney23/go-graphql-todo/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SharedApp shared application instance
var SharedApp *App

// Initialize initializes the shared app instance with
// a reference to the router, database, and config. This
// allows for reuse throughout the app through a sperate
// package; to avoid import cycles.
//
// also returns a reference to the caller for convenience
func Initialize(r *gin.Engine, db *gorm.DB, c *config.AppConfig) *App {
	SharedApp = &App{r, db, c}
	return SharedApp
}
