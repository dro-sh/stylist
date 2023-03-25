package main

import (
	"os"

	"github.com/dro-sh/stylist/backend/app"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Loads the env variables
	godotenv.Load()

	// Prints the address of our api to the console
	logrus.Info("Starting Server on http://localhost:", os.Getenv("PORT"))

	// Server router on given port
	server := app.NewServer()
	server.ListenAndServe()
}
