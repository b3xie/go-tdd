package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("shall say hello to people", func(t *testing.T) {
		const person = "Chris"
		got := Hello(person)
		want := "Hello, " + person
		assertReply(t, got, want)
	})
	t.Run("empty string must return hello world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertReply(t, got, want)
	})
}
func assertReply(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %q but got %q", want, got)
	}
}
