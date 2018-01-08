package main

import (
	"net/http"

	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
)

func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func main() {
	authenticator := auth.NewBasicAuthenticator("example.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(handle))

	log.Println("Listening...")
	http.ListenAndServe(":9090", nil)
}
