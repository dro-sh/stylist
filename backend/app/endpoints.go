package app

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/dro-sh/stylist/backend/app/lib"
)

type Message struct {
	Message string `json:"token"`
}

// Login performs login to service by login and password
func Login(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")

	data := []byte(login + password)

	res := lib.Response{ResponseWriter: w}

	token := fmt.Sprintf("%x", md5.Sum(data))
	m := Message{string(token)}
	res.SendOK(m)
}
