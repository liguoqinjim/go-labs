package mac

/*
#include "share.h"
*/
import "C"
import (
	"fmt"
	"net/url"
	"unsafe"

	"github.com/murlokswarm/app"
	"github.com/satori/go.uuid"
)

type share struct {
	id uuid.UUID
}

func newShare(s app.Share) share {
	var cvalue *C.char
	defer free(unsafe.Pointer(cvalue))

	switch v := s.Value.(type) {
	case url.URL:
		cvalue = cString(v.String())
		C.Share_URL(cvalue)

	case *url.URL:
		cvalue = cString(v.String())
		C.Share_URL(cvalue)

	default:
		cvalue = cString(fmt.Sprint(v))
		C.Share_Text(cvalue)
	}
	return share{
		id: uuid.NewV1(),
	}
}

func (s share) ID() uuid.UUID {
	return s.id
}
