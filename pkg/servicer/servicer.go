package servicer

import (
	"github.com/stephenrh/redesigned-tribble/internal/app/models/mongo"

	"github.com/stephenrh/redesigned-tribble/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

func NewService() *gin.Engine {
	r := gin.Default()
	controllers.UserRoutes(r)
	InitDb()
	return r
}

func InitDb() {
	mongo.MaxPool = 8
	mongo.PATH = "localhost"
	mongo.DBNAME = "users"
	mongo.CheckAndInitServiceConnection()
}
