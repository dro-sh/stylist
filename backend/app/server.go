package app

import (
	"net/http"
	"os"
)

// NewServer ...
func NewServer() *http.Server {
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: NewRouter(),
	}

	return server
}
