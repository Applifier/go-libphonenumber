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

// Parse parses a phone number
func Parse(number, region string) PhoneNumber {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	cRegion := C.CString(region)
	defer C.free(unsafe.Pointer(cRegion))
	res := C.parse(cNum, cRegion)
	defer C.free_phone_info(res)
	return phoneNumberInfoFromCStruct(res)
}
