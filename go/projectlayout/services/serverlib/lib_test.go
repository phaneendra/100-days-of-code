package serverlib

import "testing"

func TestHello(t *testing.T) {
	want := "serverlib hello"
	got := Hello()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
