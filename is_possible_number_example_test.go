package phonenumber_test

import (
	"fmt"
	"github.com/Applifier/go-libphonenumber"
)

func ExampleIsPossibleNumber() {
	number := "+358401231234"
	defaultRegion := "FI"
	possibleNumber := phonenumber.IsPossibleNumber(number, defaultRegion)
	fmt.Printf("The number %s is possible phone number in region %s: %v", number, defaultRegion, possibleNumber)
	// Output: The number +358401231234 is possible phone number in region FI: true
}
