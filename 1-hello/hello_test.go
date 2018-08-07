package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// tell tes t suite that this method is just a helper
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	// this is a subtest
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Faiz", "")
		want := "Hello, Faiz"

		assertCorrectMessage(t, got, want)
	})

	// this is a subtest
	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in france", func(t *testing.T) {
		got := Hello("Faiz", "French")
		want := "Bonjour, Faiz"

		assertCorrectMessage(t, got, want)
	})
}
