package driver

import (
	"fmt"
)

// statusCodeErrorHandler is used to handle error codes
func statusCodeErrorHandler(respStatusCode int, errStatusCode int, errString string) {
	if respStatusCode == errStatusCode {
		var err error
		if errString != "" {
			err = fmt.Errorf(errString)
		} else {
			err = fmt.Errorf("appigo: unexpected error occured, recieved status code %d", respStatusCode)
		}
		panic(err)
	}
}
