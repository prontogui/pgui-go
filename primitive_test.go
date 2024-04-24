package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func verifyAllFieldsAttached(t *testing.T, res Reserved, fields ...string) {
	for _, field := range fields {
		if res.findField(key.FKeyFor(field)) == nil {
			t.Errorf("Field '%s' is not attached to primitive.", field)
		}
	}
}
