package tutils

import (
	"testing"
)

func EqualS(t *testing.T, expected string, actual string, name string) {
	if expected != actual {
		t.Fatalf("%s - expected %s got %s", name, expected, actual)
	}
}

func EqualI(t *testing.T, expected int, actual int, name string) {
	if expected != actual {
		t.Fatalf("%s - expected %d got %d", name, expected, actual)
	}
}
