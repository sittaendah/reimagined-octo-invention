package main

import (
	"github.com/gin-gonic/gin"
	auth "github.com/sittaendah/aegis/internal/auth"
	"github.com/sittaendah/aegis/internal/config"
	"github.com/sittaendah/aegis/internal/mb"
	org "github.com/sittaendah/aegis/internal/organization"
	user "github.com/sittaendah/aegis/internal/user"
)

func main() {
	config.ConnectDB()

	mb.SetupKafkaProducer()
	defer mb.Producer.Close()
	r := gin.Default()

	userRepo := &user.UserRepository{DB: config.DB}
	userService := &user.UserService{Repo: userRepo}
	userController := &user.UserController{Service: userService}

	orgRepo := &org.OrganizationRepository{DB: config.DB}
	orgService := &org.OrganizationService{Repo: orgRepo, UserRepo: userRepo}
	orgController := &org.OrganizationController{Service: orgService}

	authController := &auth.AuthController{UserService: userService}

	r.POST("/login", authController.Login)

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
