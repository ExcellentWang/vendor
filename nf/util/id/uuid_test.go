package id

import (
	"testing"
)

func TestNewUUID(t *testing.T) {
	NewUUID()
}

func TestHex(t *testing.T) {
	uuidHex := NewUUID().Hex()
	if len(uuidHex) != 24 {
		t.Errorf("uuidHex length must be 24")
	}
}
