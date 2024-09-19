package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Chris")
		expected := "Hello, Chris!"
		assertCorrectMessage(t, result, expected)
	})
	t.Run("say 'Hello World!' when no name is given", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World!"
		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t testing.TB, result, expected string) {
	t.Helper()
	if (result != expected) {
		t.Errorf("got %q want %q", result, expected)
	}
}