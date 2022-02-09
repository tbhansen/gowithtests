package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Tobias")
	want := "Hello, Tobias"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
