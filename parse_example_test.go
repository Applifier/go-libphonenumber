package phonenumber_test

import (
	"fmt"
	"github.com/Applifier/go-libphonenumber"
)

func ExampleParse() {
	number := "040 123 1234"
	defaultRegion := "FI"
	info := phonenumber.Parse(number, defaultRegion)
	fmt.Printf("The number %s in normalized format is %s and validity is %v", info.Number, info.Normalized, info.Valid)
	// Output: The number 040 123 1234 in normalized format is +358401231234 and validity is true
}
