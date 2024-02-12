package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamesatomc/go-api/controllers"
	"github.com/jamesatomc/go-api/models"
)

func main() {
	server := gin.Default()

	models.ConnectDatabase()
	
	// User
	server.GET("/users", controllers.FindUsers)
	server.POST("/users", controllers.CreateUser)
	server.GET("/users/:id", controllers.FindUser)
	server.PATCH("/users/:id", controllers.UpdateUser)
	server.DELETE("/users/:id", controllers.DeleteUser)
	

	server.Run()
	// err := server.Run(":8080")
	// if err != nil {
	// 		return
	// }
}
