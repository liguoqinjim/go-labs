package mac

import "C"
import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
)

//export onJSCall
func onJSCall(cmsg *C.char) {
	msg := C.GoString(cmsg)

	app.UIChan <- func() {
		app.HandleEvent(msg)
	}
}

//export onJSAlert
func onJSAlert(calert *C.char) {
	alert := C.GoString(calert)

	app.UIChan <- func() {
		log.Warn(alert)
	}
}
