package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	env := make(chan string, 1)

	log.Printf("Starting up env: %v", os.Getenv("ENVTORUN"))
	if os.Getenv("ENVTORUN") == "prod" {
		godotenv.Load(".env.prod")
		env <- os.Getenv(key)
	} else {
		godotenv.Load(".env.dev")
		env <- os.Getenv(key)
	}

	return <-env
}
