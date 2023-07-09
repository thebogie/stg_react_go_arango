package connection

import (
	"back/infrastructure"
	"context"
	"log"
	"os"
	"sync"
	"time"

	driver "github.com/arangodb/go-driver"

	"github.com/arangodb/go-driver/http"
	"github.com/sirupsen/logrus"
)

type DatabaseConnection struct {
	Client driver.Client
	Db     driver.Database
}

var (
	instance *DatabaseConnection
	once     sync.Once
)

func GetDatabaseConnection() (*DatabaseConnection, error) {
	once.Do(func() {

		var err error

		var conn driver.Connection

		databaseURI := make(chan string, 1)

		databaseURI <- dotenv.GodotEnv("DATABASE_URI")

		log.Printf("DB CONNECTION ENVTORUN SETTING:%v:::DB URI:%v", os.Getenv("ENVTORUN"), dotenv.GodotEnv("DATABASE_URI"))

		var uptime = 10
		var client driver.Client

		conn, err = http.NewConnection(http.ConnectionConfig{
			Endpoints: []string{<-databaseURI},
		})

		if err != nil {
			logrus.Fatal("Creation of connection string to failed", err.Error())
		}

		client, err = driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication("root", "letmein"),
			//Authentication: driver.BasicAuthentication("root", "wnbGnPpCXHwbP"),
		})
		if err != nil {
			logrus.Fatal("Creation of NewClient failed", err.Error())
		}

		client.Connection().SetAuthentication(driver.BasicAuthentication("root", "letmein"))

		for i := 0; i < uptime; i++ {
			log.Printf("Attempting arangodb connection to smacktalk db")
			time.Sleep(5 * time.Second)

			ctx := context.Background()
			huh, err := client.Databases(ctx)
			logrus.Info("huh", huh)
			what, err := client.DatabaseExists(ctx, "smacktalk")
			if err != nil {
				logrus.Info("Try again. Smacktalk DB isnt found:   ", err.Error())

				if i == uptime {
					logrus.Fatal("Smacktalk DB isnt found after x times", err.Error())
				}
			} else {
				logrus.Info("Found smacktalk:", what)
				//worked
				i = 10
			}

		}

		ctx := context.Background()

		huh, err := client.Databases(ctx)
		if err != nil {
			log.Printf("ERROR:", err.Error())
		} else {

			log.Printf("huh:", huh)
		}

		users, err := client.Users(ctx)
		if err != nil {
			log.Printf("ERROR:", err.Error())
		} else {

			log.Printf("users:", users)
		}
		dbca, err := client.Databases(ctx)
		if err != nil {
			log.Printf("DBS ERROR:", err.Error())
		} else {
			log.Printf("DBS:", dbca)
		}

		db, err := client.Database(ctx, "smacktalk")

		if err != nil {
			defer logrus.Info("Connection to smacktalk Failed")
			logrus.Fatal("Smacktalk database not reachable", err.Error())
		} else {
			logrus.Info("Connection to Database Successfully")
		}

		instance = &DatabaseConnection{
			Client: client,
			Db:     db,
		}

	})
	return instance, nil
}

func Query() {

}
