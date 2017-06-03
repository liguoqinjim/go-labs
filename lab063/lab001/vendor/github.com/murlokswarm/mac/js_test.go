package mac

import "testing"
import "unsafe"

func TestOnJSCall(t *testing.T) {
	msg := cString("hello")
	defer free(unsafe.Pointer(msg))

	onJSCall(msg)
}

func TestOnJSAlert(t *testing.T) {
	alert := cString("alert")
	defer free(unsafe.Pointer(alert))

	onJSAlert(alert)
}
