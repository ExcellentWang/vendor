package nconf

import (
	"testing"
)

func TestIsEncValue(t *testing.T) {
	if isEncValue("") {
		t.Fail()
	}
	if isEncValue("hello") {
		t.Fail()
	}
	if isEncValue("ENC(a") {
		t.Fail()
	}
	if isEncValue("ENC(a") {
		t.Fail()
	}
	if isEncValue("a)") {
		t.Fail()
	}
	if !isEncValue("ENC(a)") {
		t.Fail()
	}
}

func TestGetInnerEncValue(t *testing.T) {
	if getInnerEncValue("ENC(abcd)") != "abcd" {
		t.Fail()
	}
}
