// +build !snet_trace

package snet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func (l *Listener) trace(format string, args ...interface{}) {
	//log.Printf("Listener: "+format, args...)
	log.Warnf("Listener: "+format, args...)
}

func (c *Conn) trace(format string, args ...interface{}) {
	if c.listener == nil {
		format = fmt.Sprintf("Client conn %d: %s", c.id, format)
		log.Errorf(format, args...)
	} else {
		format = fmt.Sprintf("Server conn %d: %s", c.id, format)
		log.Infof(format, args...)
	}
	//log.Printf(format, args...)

}
