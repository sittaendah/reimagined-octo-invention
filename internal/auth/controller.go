package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sittaendah/aegis/internal/mb"
	user "github.com/sittaendah/aegis/internal/user"
	"net/http"
)

type AuthController struct {
	UserService *user.UserService
}

func (c *AuthController) Login(ctx *gin.Context) {
	var credential struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := c.UserService.GetUserByUsername(credential.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if !c.UserService.CheckPassword(credential.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		mb.SendMessage(user.Username)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err, "token": token})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
