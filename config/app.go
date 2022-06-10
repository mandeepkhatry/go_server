package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppParams *AppConfig

type AppConfig struct {
	Host    string
	Port    string
	ApiBase string
}

var EnvironmentValues = struct {
	AuthToken string
}{
	"",
}

func init() {
	godotenv.Load(".env")
	EnvironmentValues.AuthToken = os.Getenv(EnvironmentKey.AuthToken)

}
