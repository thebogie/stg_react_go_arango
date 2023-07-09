package graph

import (
	"back/graph/generated"
	genmodel "back/graph/generated/model"
	"back/internal/usecase"
	"back/web/middleware/auth"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type Resolver struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

func NewResolver(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) *Resolver {
	return &Resolver{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) RegisterUser(ctx context.Context, username string, email string, password string) (*genmodel.AuthPayload, error) {

	// Invoke the corresponding use case method using the generated model types
	result, err := r.Resolver.userUsecase.RegisterUser(username, email, password)
	if err != nil {
		return nil, err
	}

	// Convert the result to the generated AuthPayload type
	authPayload := &genmodel.AuthPayload{
		Token: result.Token,
		User: &genmodel.User{
			ID:        result.User.ID,
			Firstname: result.User.Firstname,
			Email:     result.User.Email,
		},
	}

	return authPayload, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, email string, password string) (*genmodel.AuthPayload, error) {
	//TODO: get via config
	var jwtKey = []byte("your-secret-key")

	// Invoke the corresponding use case method using the generated model types
	founduser, err := r.Resolver.authUsecase.LoginUser(email, password)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"user_id": founduser.User.ID,
		"email":   founduser.User.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	//passed save cookie
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
	}

	var writeit = ctx.Value("auth-cookie").(*auth.CookieAccess).Writer
	// Set the cookie in the response
	http.SetCookie(writeit, cookie)

	// Convert the result to the generated AuthPayload type
	authPayload := &genmodel.AuthPayload{
		Token: tokenString,
		User: &genmodel.User{
			ID:        founduser.User.ID,
			Firstname: founduser.User.Firstname,
			Email:     founduser.User.Email,
		},
	}

	CA := auth.GetCookieAccess(ctx)
	CA.SetToken(tokenString)
	CA.Id = founduser.User.ID
	CA.IsLoggedIn = true

	return authPayload, nil

}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) FindUserByEmail(ctx context.Context, email string) (*genmodel.User, error) {
	// Implement the logic to retrieve the authenticated user

	CA := auth.GetCookieAccess(ctx)
	if !CA.IsLoggedIn {
		return &genmodel.User{}, fmt.Errorf("Access denied")
	}

	founduser, err := r.userUsecase.FindUserByEmail(email)

	if err != nil {

	}

	return founduser, err
}
