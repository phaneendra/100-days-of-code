package auth

import "testing"

func TestAuth(t *testing.T) {
	want := "auth library auth"
	got := Auth()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
