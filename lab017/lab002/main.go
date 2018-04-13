package main

import (
	"log"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`((https?):\/\/)?(([^:\n\r]+):([^@\n\r]+)@)?((www\.)?([^/\n\r]+))\/?([^?\n\r]+)?\??([^#\n\r]*)?#?([^\n\r]*)`)
	content := "http://someuser:password@www.website.com/path/to/file.ext?query=something&parameters=passed#with-anchor"

	r1 := r.FindString(content)
	log.Println("r1=", r1)

	r2 := r.FindStringSubmatch(content)
	for n, v := range r2 {
		log.Printf("r2[%d]=%s", n, v)
	}
}
