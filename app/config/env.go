package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {

	mongoIRU := strings.TrimSpace(os.Getenv("MONGOURI"))
	if len(mongoIRU) == 0 {

		err := godotenv.Load()

		if err != nil {
			log.Fatalln("error .env")
		}

		mongoIRU = os.Getenv("MONGOURI")
	}
	return mongoIRU
}

func GetPort() string {

	port := strings.TrimSpace(os.Getenv("PORT"))

	if len(port) == 0 {

		err := godotenv.Load()

		if err != nil {
			return "8080"

		}

		port = os.Getenv("PORT")
	}

	return port
}

func GetSecret() string {

	secrets := strings.TrimSpace(os.Getenv("SECRETKEY"))

	return secrets
}

func GetWebHookURL() string {

	secrets := strings.TrimSpace(os.Getenv("WEBHOOK_URL"))

	return secrets
}
