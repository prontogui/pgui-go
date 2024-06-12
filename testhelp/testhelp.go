// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testhelp

import (
	"testing"
)

func TestErrorMessage(t *testing.T, err error, expected string) {

	if err == nil {
		t.Fatalf("function should have returned an error")
	}

	actual := err.Error()

	if actual != expected {
		t.Fatalf("function returned the wrong error of '%v'.  Expecting '%v'", actual, expected)
	}
}

func TestNilError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("not expecting an error from function.  Error was '%v'", err.Error())
	}
}
