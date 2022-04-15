package main

import (
	"net/http"
	"rest/config"
	"rest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.InitDb()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "first APi"})
	})

	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)
	router.PUT("/users/:userid", controller.UpdateUser)
	router.GET("/users/:userid", controller.GetUserById)
	router.DELETE("/users/:userid", controller.DeleteUser)

	router.GET("/users/:userid/addresses", controller.GetAddress)
	router.POST("/users/:userid/addresses", controller.CreateAddress)
	router.PUT("/users/:userid/addresses/:addressid", controller.UpdateAddress)
	router.GET("/users/:userid/addresses/:addressid", controller.GetAddressById)
	router.DELETE("/users/:userid/addresses/:addressid", controller.DeleteAddress)

	router.Run()

}
