package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sittaendah/aegis/internal/config"
	org "github.com/sittaendah/aegis/internal/organization"
	user "github.com/sittaendah/aegis/internal/user"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	userRepo := &user.UserRepository{DB: config.DB}
	userService := &user.UserService{Repo: userRepo}
	userController := &user.UserController{Service: userService}

	orgRepo := &org.OrganizationRepository{DB: config.DB}
	orgService := &org.OrganizationService{Repo: orgRepo}
	orgController := &org.OrganizationController{Service: orgService}

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	r.POST("/organizations", orgController.CreateOrganization)
	r.GET("/organizations", orgController.GetAllOrganizations)
	r.GET("/organizations/:id", orgController.GetOrganization)
	r.PUT("/organizations/:id", orgController.UpdateOrganization)
	r.DELETE("/organizations/:id", orgController.DeleteOrganization)

	r.Run(":8080")
}
