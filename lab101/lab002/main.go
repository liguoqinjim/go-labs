package main

import (
	"expvar"
	"fmt"
	"net/http"
	"time"
)

//map
var stats = expvar.NewMap("tcp")
var requests, requestsFailed expvar.Int

//function
var start = time.Now()

func calculateUpTime() interface{} {
	return time.Since(start).String()
}

func init() {
	stats.Set("requests", &requests)
	stats.Set("requests_failed", &requestsFailed)

	expvar.Publish("upTime", expvar.Func(calculateUpTime))
}

func handler(w http.ResponseWriter, r *http.Request) {
	requests.Add(1)
	requestsFailed.Add(100)
	fmt.Fprintf(w, "Hi there,I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1818", nil)
}
