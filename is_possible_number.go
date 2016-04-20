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
	"unsafe"
)

// IsPossibleNumber can be used for quickly guessing whether a number is a possible phonenumber by using only the length information, much faster than a full validation.
func IsPossibleNumber(number, region string) bool {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	cRegion := C.CString(region)
	defer C.free(unsafe.Pointer(cRegion))
	res := C.is_possible_number(cNum, cRegion)
	return res != 0
}
