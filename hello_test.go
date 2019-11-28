package gomod

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World!"
	if Hello() != want {
		t.Errorf("Hello() != %s", want)
	}
}
