// internal/service/user_service.go
package service

import (
	genmodel "back/graph/generated/model"
	repo "back/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(username, email, password string) (*genmodel.AuthPayload, error)
	FindUserByEmail(email string) (*genmodel.User, error)
	FindUserByID(userid string) (*genmodel.User, error)
}

type userService struct {
	userRepository repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Implement the UserService methods
func (s *userService) RegisterUser(username, email, password string) (*genmodel.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}

func (s *userService) FindUserByEmail(email string) (*genmodel.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}

func (s *userService) FindUserByID(id string) (*genmodel.User, error) {
	user, err := s.userRepository.FindUserByID(id)
	logrus.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}
