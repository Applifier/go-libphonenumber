package phonenumber

import (
	"encoding/json"
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
	if *res.Normalized != "+358401231234" {
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
	if *res.Normalized != "+358401231234" {
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
	if res.Normalized != nil {
		t.Error("Should not normalize number")
	}
	if res.Error == nil {
		t.Error("Did not get error when expected")
	}
	runtime.GC()
}

func TestJSONMarshal(t *testing.T) {
	info := PhoneNumber{
		Valid:      false,
		Error:      NewError("decoding failed"),
		Number:     "abadogjwprj",
		Normalized: nil,
	}
	data, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	parsedInfo := PhoneNumber{}
	err = json.Unmarshal(data, &parsedInfo)
	if err != nil {
		t.Fatal(err)
	}
	if info.Valid != parsedInfo.Valid || info.Error.Error() != parsedInfo.Error.Error() || info.Number != parsedInfo.Number || parsedInfo.Normalized != nil {
		t.Error("Marshaled and unmarshaled structs don't match.")
	}

	normalized := "+358401231234"
	info = PhoneNumber{
		Valid:      true,
		Error:      nil,
		Number:     "+358 40 123 1234",
		Normalized: &normalized,
	}
	data, err = json.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	parsedInfo = PhoneNumber{}
	err = json.Unmarshal(data, &parsedInfo)
	if err != nil {
		t.Fatal(err)
	}
	if info.Valid != parsedInfo.Valid || parsedInfo.Error != nil || info.Number != parsedInfo.Number || *info.Normalized != *parsedInfo.Normalized {
		t.Error("Marshaled and unmarshaled structs don't match.")
	}
}
