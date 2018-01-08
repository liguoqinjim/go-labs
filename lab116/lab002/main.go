package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
)

func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "b98e16cbc3d01734b264adba7baa3bf9"
	}
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func main() {
	authenticator := auth.NewDigestAuthenticator("example.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(handle))

	log.Println("Listening...")
	http.ListenAndServe(":9090", nil)
}
