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

// CountryCode returns the country code for the supplied number
func CountryCode(number string) string {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	res := C.get_country_code(cNum)
	defer C.free(unsafe.Pointer(res))
	return C.GoString(res)
}
