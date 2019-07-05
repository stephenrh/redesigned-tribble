package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenrh/gosvc1/internal/app/models"
)

func index(c *gin.Context) {
	var (
		users []models.User
		err   error
	)
	users, err = models.GetUsers()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.JSON(500, errResp)
		return
	}
	c.JSON(200, users)
}

func create(c *gin.Context) {
	var user models.User
	var err error
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(400, err)
		return
	}
	ruser, err := models.CreateUser(user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ruser)
	return
}

func show(c *gin.Context) {
	id := c.Param("id")
	c.String(200, id)
	return
}

func UserRoutes(r *gin.Engine) {
	users := r.Group("/")
	{
		users.GET("/", index)
		users.GET("/:id", show)
		users.POST("/", create)
	}
}
