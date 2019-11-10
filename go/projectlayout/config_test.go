package projectlayout

import "testing"

func TestConfig(t *testing.T) {
	want := "projectlayout config"
	got := Config()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
