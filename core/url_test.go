package core

import (
	"testing"
)

func TestFixProtoPrefix(t *testing.T) {
	original := "www.google.com/"
	expected := "https://www.google.com/"
	res := FormatUri(original)
	if res.Modified != expected {
		t.Errorf("Expected %s got %s", expected, original)
	}
}

func TestIsModified(t *testing.T) {
	res := FormatUri("something")
	res.Modified = "somethingelse"
	if !res.IsModified() {
		t.Errorf("Uri was modified but IsModified returned false")
	}
}
