package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// a new instance of the logger
var log = logrus.New()

func main() {
	log.Out = os.Stdout

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
