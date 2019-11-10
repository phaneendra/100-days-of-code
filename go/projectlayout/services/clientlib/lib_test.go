package clientlib

import "testing"

func TestHello(t *testing.T) {
	want := "clientlib hello"
	got := Hello()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
