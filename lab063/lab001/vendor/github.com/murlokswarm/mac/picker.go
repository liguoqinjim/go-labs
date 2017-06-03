package mac

/*
#include "picker.h"
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
	"github.com/satori/go.uuid"
)

type filePicker struct {
	id     uuid.UUID
	picker app.FilePicker
}

func newFilePicker(p app.FilePicker) *filePicker {
	id := uuid.NewV1()

	cpicker := C.FilePicker__{
		ID:                cString(id.String()),
		MultipleSelection: boolToBOOL(p.MultipleSelection),
		NoDir:             boolToBOOL(p.NoDir),
		NoFile:            boolToBOOL(p.NoFile),
	}
	defer free(unsafe.Pointer(cpicker.ID))

	picker := &filePicker{
		id:     id,
		picker: p,
	}
	app.Elements().Add(picker)

	C.Picker_NewFilePicker(cpicker)
	return picker
}

func (p *filePicker) ID() uuid.UUID {
	return p.id
}

//export onFilePickerClosed
func onFilePickerClosed(cid *C.char, filenamesJSON *C.char) {
	id := uuid.FromStringOrNil(C.GoString(cid))
	elem, ok := app.Elements().Get(id)
	if !ok {
		return
	}
	defer app.Elements().Remove(elem)

	var filenames []string
	data := []byte(C.GoString(filenamesJSON))
	if err := json.Unmarshal(data, &filenames); err != nil {
		log.Error(err)
		return
	}

	p := elem.(*filePicker)
	if len(filenames) != 0 && p.picker.OnPick != nil {
		app.UIChan <- func() { p.picker.OnPick(filenames) }
	}
}
