package tutils

import (
	"reflect"
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

func EqualNil(t *testing.T, obj interface{}, name string) {
	if obj == nil {
		return
	}
	if !reflect.ValueOf(obj).IsNil() {
		t.Fatalf("%s - expected %#v to be nil but it wasnt", name, obj)
	}
}

func EqualNotNil(t *testing.T, obj interface{}, name string) {
	if reflect.ValueOf(obj).IsNil() {
		t.Fatalf("%s - expected %#v to not be nil but it wasnt", name, obj)
	}
}

func EqualB(t *testing.T, expected bool, actual bool, name string) {
	if expected != actual {
		t.Fatalf("%s - expected %t got %t", name, expected, actual)
	}
}
