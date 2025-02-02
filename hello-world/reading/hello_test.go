package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("given a name, expect greeting", func(t *testing.T) {
		got := Hello("Alice", "")
		want := "Hello, Alice"
		assertCorrectMessage(t, got, want)
	})
	t.Run("given empty name, expect default greeting", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
