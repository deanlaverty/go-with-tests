package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Return hello dean when name is passed", func(t *testing.T) {
		got := Hello("Dean", "English")
		want := "Hello, Dean"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return hello world when no string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return name in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return name in french", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Hello in French, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
