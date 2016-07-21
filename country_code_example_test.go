package phonenumber_test

import (
	"fmt"

	"github.com/Applifier/go-libphonenumber"
)

func ExampleCountryCodeFi() {
	number1 := "+358401231234"
	cc1 := phonenumber.CountryCode(number1)
	number2 := "+15417543010"
	cc2 := phonenumber.CountryCode(number2)
	fmt.Printf("%s and %s", cc1, cc2)
	// Output: FI and US
}
