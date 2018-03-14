package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port = "8000"

func main() {
	http.HandleFunc("/hi", hi)

	fmt.Println("runing on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hi(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "HostName: %s", hostName)
}
