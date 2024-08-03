package auth

import (
	"fmt"

	user "github.com/sittaendah/aegis/internal/user"
)

type AuthServiceInterface interface {
	Login(username string, password string) bool
}

type AuthService struct {
	user.UserRepository
	user.UserService
}

type AuthCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s AuthService) Login(username string, password string) bool {
	var user user.User
	user, err := s.UserRepository.GetUserByUsername(username)
	fmt.Println(user)
	if err != nil {
		fmt.Errorf("Login: %v", err)
		return false
	}
	return s.UserService.CheckPassword(password, user.Password)
}
