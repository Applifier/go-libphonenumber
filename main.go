package main

/*
#cgo CPPFLAGS: -I/usr/local/include
#cgo CPPFLAGS: -Wall -Werror -std=c++14 -x c++
#cgo LDFLAGS: -llibphonenumber
#include <phonenumbers/phonenumberutil.h>
using namespace i18n::phonenumbers;
*/
import "C"

import (
	"log"
)

func main() {
	cNum := C.CString("+358407597575")
	res := C.isPossibleNumber(cNum)
	log.Printlnf("Response %v", res)
}
