package web

import (
	"back/internal/usecase"
	"back/web/middleware/auth"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"io"
	"net/http"
	"strings"

	"back/graph"
	"back/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func NewServer(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) http.Handler {
	router := chi.NewRouter()

	router.Use(Logger)
	router.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Use(auth.Middleware(userUsecase))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(userUsecase, authUsecase)}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	return router
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("**Request:")
		fmt.Println(formatRequest(r))

		wrapped := wrapResponseWriter(w)
		fmt.Println("**END OF REQUEST**")
		next.ServeHTTP(wrapped, r)

		fmt.Println("**Response:")
		fmt.Println("Status:", wrapped.Status())
		fmt.Println("Header:", wrapped.Header())
		fmt.Println("Body:", wrapped.Body.String())
		fmt.Println("**END OF RESPONSE**")
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	StatusNum int
	Body      strings.Builder
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{
		ResponseWriter: w,
		StatusNum:      http.StatusOK,
	}
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.StatusNum = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWrapper) Status() int {
	return rw.StatusNum
}

func formatRequest(r *http.Request) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Method: %s\n", r.Method))
	sb.WriteString(fmt.Sprintf("URL: %s\n", r.URL.String()))
	sb.WriteString("Headers:\n")
	r.Header.Write(&sb)
	sb.WriteString("\n")

	// Copy the original body to a new buffer
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		sb.WriteString(fmt.Sprintf("Failed to read request body: %s\n", err.Error()))
	} else {
		sb.WriteString(fmt.Sprintf("Body: %s\n", buf.String()))
	}

	// Restore the original body
	r.Body = io.NopCloser(strings.NewReader(buf.String()))

	return sb.String()
}
