package app

import (
	"net/http"

	"github.com/dro-sh/stylist/backend/app/lib"
)

// Message ...
type Message struct {
	Message string `json:"message"`
}

// HelloWorld ...
func Login(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}

	m := Message{"Hello User!"}
	res.SendOK(m)
}
