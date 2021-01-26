package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jeansferreira/api-b2w-planetas/environment"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	Host     string
	Database string
	Password string
	Port     string
	User     string
	db       *mgo.Database
}

func (m *MongoDB) ConnectionUrl() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		m.User, m.Password, m.Host, m.Port, m.Database,
	)
}

func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.ConnectionUrl())
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.Database)
}

func getFromEnvIfExists() MongoDB {

	host, envHost := os.LookupEnv("DB_HOST")
	database, envDatabase := os.LookupEnv("DB_NAME")
	password, envPassword := os.LookupEnv("DB_PASSWORD")
	user, envUser := os.LookupEnv("DB_USER")
	port, envPort := os.LookupEnv("DB_PORT")

	if !envHost || !envDatabase || !envPassword || !envUser || !envPort {
		return MongoDB{
			Host:     environment.DATABASE_HOST,
			Database: environment.DATABASE_NAME,
			Password: environment.DATABASE_PASSWORD,
			User:     environment.DATABASE_USER,
			Port:     environment.DATABASE_PORT,
		}
	}

	newMongoDB := MongoDB{Host: host, Database: database, Password: password, User: user, Port: port}

	return newMongoDB
}

var Mongo = getFromEnvIfExists()
