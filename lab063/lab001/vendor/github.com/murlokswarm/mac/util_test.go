package mac

import "testing"
import "unsafe"

func TestBoolToBOOL(t *testing.T) {
	if boolToBOOL(true) != 1 {
		t.Error("boolToBOOL should return 1")
	}

	if boolToBOOL(false) != 0 {
		t.Error("boolToBOOL should return 0")
	}
}

func TestCString(t *testing.T) {
	cstr := cString("Maxence")
	free(unsafe.Pointer(cstr))
}
