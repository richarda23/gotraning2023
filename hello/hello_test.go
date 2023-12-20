package main

import (
	"hello/testhelper"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brian", "")
		wanted := "Hello, Brian"
		testhelper.AssertCorrectMessage(t, got, wanted)
	})
	t.Run("general hllo", func(t *testing.T) {
		got := Hello("", "")
		wanted := "Hello, World"
		testhelper.AssertCorrectMessage(t, got, wanted)
	})
	t.Run("hello in spanish", func(t *testing.T) {
		got := Hello("", "spanish")
		wanted := "Hola, World"
		testhelper.AssertCorrectMessage(t, got, wanted)
	})

	t.Run("hello in french", func(t *testing.T) {
		got := Hello("emilie", "french")
		wanted := "Bonjour, emilie"
		testhelper.AssertCorrectMessage(t, got, wanted)
	})
}
