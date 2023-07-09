// internal/usecase/user_usecase.go
package usecase

import (
	genmodel "back/graph/generated/model"
	service "back/internal/service"
	"log"
)

type UserUsecase interface {
	RegisterUser(username, email, password string) (*genmodel.AuthPayload, error)
	FindUserByEmail(email string) (*genmodel.User, error)
	FindUserByID(id string) (*genmodel.User, error)
}

type userUsecase struct {
	userService service.UserService
}

func NewUserUsecase(userService service.UserService) UserUsecase {
	return &userUsecase{
		userService: userService,
	}
}

// Implement the UserUsecase methods
func (u *userUsecase) RegisterUser(username, email, password string) (*genmodel.AuthPayload, error) {
	// Implement the logic to register a user
	return nil, nil
}

func (u *userUsecase) FindUserByEmail(email string) (*genmodel.User, error) {
	user, err := u.userService.FindUserByEmail(email)
	log.Printf(user.Email)
	if err != nil {
	}
	return user, nil
}

func (u *userUsecase) FindUserByID(userid string) (*genmodel.User, error) {
	user, err := u.userService.FindUserByID(userid)
	log.Printf(user.Email)
	if err != nil {
	}
	return nil, nil
}
