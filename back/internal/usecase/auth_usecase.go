// internal/usecase/auth_usecase.go
package usecase

import (
	genmodel "back/graph/generated/model"
	service "back/internal/service"
)

type AuthUsecase interface {
	LoginUser(email, password string) (*genmodel.AuthPayload, error)
	AuthUser(email, password string) (*genmodel.AuthPayload, error)
}

type authUsecase struct {
	authService service.AuthService
}

func NewAuthUsecase(authService service.AuthService) AuthUsecase {
	return &authUsecase{
		authService: authService,
	}
}

// Implement the authUsecase methods
func (u *authUsecase) LoginUser(email, password string) (*genmodel.AuthPayload, error) {
	pay, err := u.authService.LoginUser(email, password)
	if err != nil {
		//error
	}

	return pay, err
}

func (u *authUsecase) AuthUser(email, password string) (*genmodel.AuthPayload, error) {
	// Implement the logic to register a user
	auth, err := u.authService.CheckAuthUser(email, password)
	if err != nil {

	}
	return auth, nil
}
