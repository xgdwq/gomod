package gomod

import "testing"

func TestHello(t *testing.T) {
	want := "v2: Hello World!"
	if Hello() != want {
		t.Errorf("Hello() != %s", want)
	}
}
