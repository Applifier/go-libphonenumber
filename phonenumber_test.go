package phonenumber

import (
	"testing"
)

func TestIsPossiblePhoneNumber(t *testing.T) {
	num := "+358401231234"
	region := "FI"
	if !IsPossibleNumber(num, region) {
		t.Error(num, "was not valid")
	}
	num = "asd"
	if IsPossibleNumber(num, region) {
		t.Error(num, "was valid")
	}
	num = ""
	if IsPossibleNumber(num, region) {
		t.Error(num, "was valid")
	}
	num = "0401231234"
	if !IsPossibleNumber(num, region) {
		t.Error(num, "was not valid in", region)
	}
}
