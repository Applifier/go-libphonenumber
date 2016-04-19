package phonenumber

import (
	"runtime"
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
	runtime.GC()
}

func TestParse(t *testing.T) {
	res := Parse("+358401231234", "FI")
	if !res.Valid {
		t.Error("Was not valid")
	}
	if res.Number != "+358401231234" {
		t.Error("Did not set original number")
	}
	if res.Normalized != "+358401231234" {
		t.Error("Did not normalize number")
	}
	if res.Error != nil {
		t.Error("Got unexpected erro", res.Error)
	}
	runtime.GC()
	res = Parse("+358-40-123 1234", "FI")
	if !res.Valid {
		t.Error("Was not valid")
	}
	if res.Number != "+358-40-123 1234" {
		t.Error("Did not set original number")
	}
	if res.Normalized != "+358401231234" {
		t.Error("Did not normalize number")
	}
	if res.Error != nil {
		t.Error("Got unexpected erro", res.Error)
	}
	runtime.GC()
	res = Parse("acdegwerigh qepgj fqwpeg", "FI")
	if res.Valid {
		t.Error("Was valid when should not")
	}
	if res.Number != "acdegwerigh qepgj fqwpeg" {
		t.Error("Did not set original number")
	}
	if res.Normalized != "" {
		t.Error("Should not normalize number")
	}
	if res.Error == nil {
		t.Error("Did not get error when expected")
	}
	runtime.GC()
}
