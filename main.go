package main

import (
	"github.com/gin-gonic/gin"
	"reviewtest/controller"
)

func v1Routes(r *gin.RouterGroup) {
	r.GET("/ping", controllers.PingV1)
	r.POST("/validate", controllers.ValidatePasswordV1)
	r.GET("/hash", controllers.GetHashV1)
}


func defaultResponse(c *gin.Context) {
	defaultJSON := gin.H{
		"task1": "Create an endpoint which returns a 200 response with \"pong\" as response.",
		"task2": "Create an endpoint which takes a password as input and return the strenth of password.",
		"task3": "Create an endpoint which returns MD5 Hash of any input which is passed.",
	}

	c.JSON(200, defaultJSON)
}

func main() {
	r := gin.Default()
	r.GET("/", defaultResponse)
	apiV1 := r.Group("api/v1")

	v1Routes(apiV1)
	

	// RUNNING ON DEFAULT PORT 8080
	r.Run()
}
