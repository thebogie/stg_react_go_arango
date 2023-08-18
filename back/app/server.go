package main

import (
	"github.com/arangodb/go-driver"
	"log"
	"net/http"
	"os"

	"back/app/db"
	"back/auth"
	"back/graph/resolver"
	"back/pkg/adapter/controller"
	"back/pkg/registry"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "50002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	initdb := db.InitDB()

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	ctrl := newController(initdb)
	srv := handler.NewDefaultServer(resolver.NewSchema(ctrl))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func newController(db driver.Database) controller.Controller {
	r := registry.New(db)
	return r.NewController()
}
