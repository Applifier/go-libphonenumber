package phonenumber

/*
#cgo CPPFLAGS: -I/usr/local/include
#cgo CPPFLAGS: -Wall -Werror
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -lphonenumber

#include <stdlib.h>
#include "phonenumber.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// IsPossibleNumber quickly guessing whether a number is a possible phonenumber by using only the length information, much faster than a full validation.
func IsPossibleNumber(number, region string) bool {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	cRegion := C.CString(region)
	defer C.free(unsafe.Pointer(cRegion))
	res := C.isPossibleNumber(cNum, cRegion)
	return res == 1
}

// cStr returns a pointer to the first byte in s.
func cStr(s string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(h.Data))
}
