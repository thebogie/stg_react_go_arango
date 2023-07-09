package auth

import (
	"back/internal/usecase"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var cookieName = "auth-cookie"

type CookieAccess struct {
	Writer     http.ResponseWriter
	Id         string
	IsLoggedIn bool
}

// method to write cookie
func (this *CookieAccess) SetToken(token string) {
	http.SetCookie(this.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(1),
	})
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(u usecase.UserUsecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie(cookieName)
			//TODO: get via config
			var jwtKey = []byte("your-secret-key")
			//setup cookie placeholder
			cookieA := CookieAccess{
				Writer:     w,
				Id:         "",
				IsLoggedIn: false,
			}

			ctx := context.WithValue(r.Context(), cookieName, &cookieA)

			if err != nil || c == nil {
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// Parse the JWT token
			token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
				// Provide the secret key used for signing the token
				return jwtKey, nil
			})

			if err != nil {
				fmt.Println("Error parsing JWT:", err)
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// Verify if the token is valid
			if !token.Valid {
				fmt.Println("Invalid token")
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			// Extract the user_id claim from the token

			claims := token.Claims.(jwt.MapClaims)

			// put it in context
			cookieA = CookieAccess{
				Writer:     w,
				Id:         claims["user_id"].(string),
				IsLoggedIn: true,
			}

			ctx = context.WithValue(r.Context(), cookieName, &cookieA)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.

func GetCookieAccess(ctx context.Context) *CookieAccess {
	return ctx.Value(cookieName).(*CookieAccess)
}
