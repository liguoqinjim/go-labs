package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {
	numCPUs := runtime.NumCPU()
	log.Println("numCPUs:", numCPUs)

	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload
		time.Sleep(time.Second * 2)

		return result
	})
	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result := pool.Process(input)

		w.Write(result.([]byte))
	})

	http.ListenAndServe(":8080", nil)
}
