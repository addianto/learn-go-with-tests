package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("given a name, expect greeting", func(t *testing.T) {
		got := Hello("Alice")
		want := "Hello, Alice"
		if got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})
	t.Run("given empty name, expect default greeting", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got %q; want %q", got, want)
		}
	})
}
