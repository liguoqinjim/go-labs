package mac

/*
#import <Foundation/Foundation.h>
*/
import "C"
import "unsafe"

func boolToBOOL(b bool) C.BOOL {
	if b {
		return C.YES
	}
	return C.NO
}

func cString(str string) *C.char {
	return C.CString(str)
}

func free(p unsafe.Pointer) {
	C.free(p)
}
