// Package mac implements the macOS driver.
// Usage:
// import _ "github.com/murlokswarm/mac"
// During initialization, the package calls yui.RegisterDriver() with its
// Driver implementation.
package mac

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework Cocoa
#cgo LDFLAGS: -framework WebKit
#cgo LDFLAGS: -framework CoreImage
#cgo LDFLAGS: -framework Security
#include "driver.h"
*/
import "C"
import (
	"encoding/json"
	"net/url"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
	"github.com/pkg/errors"
)

var (
	driver *Driver
)

func init() {
	driver = NewDriver()
	app.RegisterDriver(driver)
}

// Driver is the implementation of the MacOS driver.
type Driver struct {
	appMenu app.Contexter
	dock    app.Docker
	running bool
}

// NewDriver creates a new MacOS driver.
// It initializes the Cocoa app.
func NewDriver() *Driver {
	return &Driver{
		appMenu: newMenuBar(),
		dock:    newDock(),
	}
}

// Run launches the Cocoa app.
func (d *Driver) Run() {
	d.running = true
	C.Driver_Run()
}

// NewElement creates a new app element.
func (d *Driver) NewElement(e interface{}) app.Elementer {
	driver.mustRun()

	switch elem := e.(type) {
	case app.Window:
		return newWindow(elem)

	case app.ContextMenu:
		return newContextMenu(elem)

	case app.Share:
		return newShare(elem)

	case app.FilePicker:
		return newFilePicker(elem)

	default:
		log.Panicf("element described by %T is not implemented", elem)
		return nil
	}
}

// MenuBar returns the menu bar.
func (d *Driver) MenuBar() (menu app.Contexter, ok bool) {
	return d.appMenu, true
}

// Dock returns the dock.
func (d *Driver) Dock() (dock app.Docker, ok bool) {
	return d.dock, true
}

// Resources returns the location of the resources directory.
func (d *Driver) Resources() string {
	return resources()
}

// Storage returns the location of the app storage directory.
func (d *Driver) Storage() string {
	return storage()
}

// JavascriptBridge returns the javascript statement to allow javascript to
// call go component methods.
func (d *Driver) JavascriptBridge() string {
	return "window.webkit.messageHandlers.Call.postMessage(msg);"
}

func (d *Driver) terminate() {
	C.Driver_Terminate()
}

func (d *Driver) mustRun() {
	if !d.running {
		log.Panic(`app is not running`)
	}
}

//export onLaunch
func onLaunch() {
	app.UIChan <- func() {
		if app.OnLaunch != nil {
			app.OnLaunch()
		}
	}
}

//export onFocus
func onFocus() {
	app.UIChan <- func() {
		if app.OnFocus != nil {
			app.OnFocus()
		}
	}
}

//export onBlur
func onBlur() {
	app.UIChan <- func() {
		if app.OnBlur != nil {
			app.OnBlur()
		}
	}
}

//export onReopen
func onReopen() {
	app.UIChan <- func() {
		if app.OnReopen != nil {
			app.OnReopen()
		}
	}
}

//export onFilesOpen
func onFilesOpen(cfilenamesJSON *C.char) {
	filenamesJSON := C.GoString(cfilenamesJSON)

	var filenames []string
	if err := json.Unmarshal([]byte(filenamesJSON), &filenames); err != nil {
		log.Error(err)
		return
	}

	app.UIChan <- func() {
		if app.OnFilesOpen != nil {
			app.OnFilesOpen(filenames)
		}
	}
}

//export onURLOpen
func onURLOpen(curl *C.char) {
	urlString := C.GoString(curl)
	URL, err := url.Parse(urlString)
	if err != nil {
		log.Error(errors.Wrap(err, "onURLOpen failed"))
		return
	}

	app.UIChan <- func() {
		if app.OnURLOpen != nil {
			app.OnURLOpen(*URL)
		}
	}
}

//export onTerminate
func onTerminate() bool {
	termChan := make(chan bool)

	app.UIChan <- func() {
		if app.OnTerminate != nil {
			termChan <- app.OnTerminate()
			return
		}

		termChan <- true
	}
	return <-termChan
}

//export onFinalize
func onFinalize() {
	if app.OnFinalize != nil {
		app.OnFinalize()
	}
}
