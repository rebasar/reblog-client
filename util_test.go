package main

import (
	"errors"
	"testing"
)

func TestPanicOnNonNilError(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("checkError should have panicked but it didn't")
		}
	}()
	func() {
		checkError(errors.New("Hello Error!"))
	}()
}

func TestDontPanicOnNilError(t *testing.T) {
	checkError(nil)
}
