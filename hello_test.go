package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brian", "")
		wanted := "Hello, Brian"
		AssertCorrectMessage(t, got, wanted)
	})
	t.Run("general hllo", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"
		AssertCorrectMessage(t, got, wanted)
	})
	t.Run("hello in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		wanted := "Hola, World"
		AssertCorrectMessage(t, got, wanted)
	})

	t.Run("hello in french", func(t *testing.T) {
		got := Hello("emilie", "french")
		wanted := "Bonjour, emilie"
		AssertCorrectMessage(t, got, wanted)
	})
}

func AssertCorrectMessage(t testing.TB, got, wanted string) {
	t.Helper()
	if got != wanted {
		t.Errorf("wanted %q but got %q", wanted, got)
	}
}
