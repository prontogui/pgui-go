package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func verifyBsideFieldsAttached(t *testing.T, bside *BSide, res Reserved) {
	if res.bside != bside {
		t.Error("BSide fields are not attached")
	}
}

func verifyAllFieldsAttached(t *testing.T, res Reserved, fields ...string) {
	for _, field := range fields {
		if res.findField(key.FKeyFor(field)) == nil {
			t.Errorf("Field '%s' is not attached to primitive.", field)
		}
	}
}
