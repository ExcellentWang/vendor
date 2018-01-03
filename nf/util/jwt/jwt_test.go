package jwt

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	token, err := GenToken("thSubject", time.Now().Add(2*time.Minute))
	if err != nil {
		t.Fail()
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Fail()
	}
	subject, err := claims.GetSubject()
	if err != nil {
		t.Fail()
	}
	if subject != "thSubject" {
		t.Fail()
	}
}

func TestParseToken(t *testing.T) {
	claims, err := ParseToken("")
	if err == nil || claims != nil {
		t.Fail()
	}
}
