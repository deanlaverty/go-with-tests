package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "This is a test"}

	got := Search(dictionary, "test")
	want := "This is a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
