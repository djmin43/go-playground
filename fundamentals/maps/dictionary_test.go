package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})

}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(want, got) {
		t.Errorf("got error %q want %q", got, want)
	}
}
