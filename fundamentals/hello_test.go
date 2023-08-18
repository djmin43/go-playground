package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Error("got %q want %q", got, want)
	}

}
