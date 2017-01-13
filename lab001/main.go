package main

import (
	"flag"
	"fmt"
)

//Example 1
var fileName = flag.String("fileName", "a.png", "config file name")

//Example 2
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

//Example 3

func main() {
	flag.Parse()

	//Example 1
	fmt.Println("fileName =", *fileName)

	//Example 2
	fmt.Println("gopherType =", gopherType)
}
