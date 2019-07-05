package servicer

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenrh/gosvc1/internal/app/controllers"
	"github.com/stephenrh/gosvc1/internal/app/models/mongo"
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
