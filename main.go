package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sittaendah/aegis/internal/config"
	user "github.com/sittaendah/aegis/internal/user"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	userRepo := &user.UserRepository{DB: config.DB}
	userService := &user.UserService{Repo: userRepo}
	userController := &user.UserController{Service: userService}

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	r.Run(":8080")
}
