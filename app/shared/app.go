package shared

import (
	"github.com/charliekenney23/go-graphql-todo/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// App type provides an interface for a shared instance
// of the applications main resources for reuse.
type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config *config.AppConfig
}
