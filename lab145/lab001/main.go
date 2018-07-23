package main

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)

	v, _ := c.Get("foo")
	log.Println("v=", v)
}
