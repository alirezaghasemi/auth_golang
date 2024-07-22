package routes

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Register
	r.POST("/register", handlers.Register)

	// Login
	r.POST("/login", handlers.Login)

	// Get All Users
	r.GET("/users", handlers.GetUsers)

	// Get User By Username
	r.GET("/users/:username", handlers.GetUserByUsername)

	// Update User By Username
	r.PUT("/users/:username", handlers.UpdateUser)

	// Delete User By Username
	r.DELETE("/users/:username", handlers.DeleteUser)

	return r
}
