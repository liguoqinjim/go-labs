package mac

/*
#include "driver.h"
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/errors"
	"github.com/murlokswarm/log"
)

type dock struct {
	*menu
}

func newDock() *dock {
	return &dock{
		menu: newMenu(app.Menu{}),
	}
}

func (d *dock) Mount(c app.Componer) {
	driver.mustRun()
	d.menu.Mount(c)
	C.Driver_SetDockMenu(d.ptr)
}

func (d *dock) Component() app.Componer {
	return d.component
}

func (d *dock) SetIcon(path string) {
	driver.mustRun()

	cpath := C.CString(path)
	defer free(unsafe.Pointer(cpath))

	if len(path) == 0 {
		C.Driver_SetDockIcon(cpath)
		return
	}

	if !app.FileIsSupportedIcon(path) {
		log.Error(errors.Newf("extension of %v is not supported", path))
		return
	}

	if _, err := os.Stat(path); err != nil {
		log.Error(errors.New(err))
		return
	}

	C.Driver_SetDockIcon(cpath)
}

func (d *dock) SetBadge(v interface{}) {
	driver.mustRun()

	if v == nil {
		v = ""
	}
	cv := C.CString(fmt.Sprint(v))
	defer free(unsafe.Pointer(cv))
	C.Driver_SetDockBadge(cv)
}
