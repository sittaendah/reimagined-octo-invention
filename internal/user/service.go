package user

import (
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *UserRepository
}

func (s *UserService) CreateUser(user User) (int, error) {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.Repo.GetUser(id)
}

func (s *UserService) GetUserByUsername(username string) (User, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) UpdateUser(user User) error {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}

func (s *UserService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
