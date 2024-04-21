package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("James")
	want := "Hello, James"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
