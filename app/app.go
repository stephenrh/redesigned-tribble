package app

import (
	"selfscale/users/app/controllers"
	"selfscale/users/app/models/mongo"

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
