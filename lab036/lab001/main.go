package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
