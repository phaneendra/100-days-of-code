package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got '%s' want '%s", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"
		assertMessage(t, got, want)

	})

	t.Run("saying hello world if no arg is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertMessage(t, got, want)
	})
}
