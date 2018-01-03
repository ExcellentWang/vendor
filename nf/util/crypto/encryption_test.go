package crypto

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	data := "Hello"
	want := "poSg-4s="
	result := Encrypt(data)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}

func TestDecrypt(t *testing.T) {
	data := "poSg-4s="
	want := "Hello"
	result := Decrypt(data)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}
