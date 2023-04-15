package main

import (
	"os"

	"github.com/arifmr/movie-festival-app/internal"
	"github.com/arifmr/movie-festival-app/internal/infra"
	"github.com/arifmr/movie-festival-app/internal/service/insert_movie"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	// Read dotenv file
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

	// Configs
	typapp.Provide("", infra.LoadPgDatabaseCfg)
	typapp.Provide("", infra.LoadEchoCfg)

	// Infra
	typapp.Provide("", infra.NewDatabases)
	typapp.Provide("", infra.NewEcho)

	// Services
	typapp.Provide("", insert_movie.New)

	if err := typapp.StartApp(internal.Start, internal.Shutdown); err != nil {
		logrus.Fatal(err.Error())
	}
}
