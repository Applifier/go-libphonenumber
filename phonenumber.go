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
	"errors"
	"unsafe"
)

// PhoneNumber stores information about a number
type PhoneNumber struct {
	Number     string `json:"number"`
	Normalized string `json:"normalized"`
	Error      error  `json:"error"`
	Valid      bool   `json:"valid"`
}

func phoneNumberInfoFromCStruct(info *C.struct_phone_info) PhoneNumber {
	pn := PhoneNumber{}
	if info == nil {
		return pn
	}
	pn.Valid = info.valid != 0
	if info.error != nil {
		pn.Error = errors.New(C.GoString(info.error))
	}
	if info.number != nil {
		pn.Number = C.GoString(info.number)
	}
	if info.normalized != nil {
		pn.Normalized = C.GoString(info.normalized)
	}

	return pn
}

// IsPossibleNumber quickly guessing whether a number is a possible phonenumber by using only the length information, much faster than a full validation.
func IsPossibleNumber(number, region string) bool {
	cNum := C.CString(number)
	defer C.free(unsafe.Pointer(cNum))
	cRegion := C.CString(region)
	defer C.free(unsafe.Pointer(cRegion))
	res := C.is_possible_number(cNum, cRegion)
	return res != 0
}

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
