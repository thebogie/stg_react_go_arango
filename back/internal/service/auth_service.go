// internal/service/auth_service.go
package service

import (
	genmodel "back/graph/generated/model"

	repo "back/internal/repository"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginUser(email, password string) (*genmodel.AuthPayload, error)
	CheckAuthUser(email, password string) (*genmodel.AuthPayload, error)
}

type authService struct {
	authRepository repo.AuthRepository
}

func NewAuthService(authRepository repo.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepository,
	}
}

// Implement the UserService methods
func (s *authService) LoginUser(email, password string) (*genmodel.AuthPayload, error) {
	pay, err := s.CheckAuthUser(email, password)
	if err != nil {
		//error
	}
	return pay, err
}

func (s *authService) CheckAuthUser(email, password string) (*genmodel.AuthPayload, error) {
	founduser, err := s.authRepository.FindByEmail(email)
	if err != nil {

	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)

	}

	fmt.Println("Hashed password:", hashedPassword)

	// Compare the password with its hash
	err = ComparePassword(password, hashedPassword)
	if err != nil {
		fmt.Println("Invalid password")

	}

	fmt.Println("Password is valid")
	logrus.Debug(founduser)

	return &genmodel.AuthPayload{
		Token: "",
		User: &genmodel.User{
			ID:        founduser.ID,
			Firstname: founduser.Firstname,
			Email:     founduser.Email,
		},
	}, err
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a salt with default cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
func ComparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
