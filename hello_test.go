package modulestut

import "testing"


func TestV3Hello(t *testing.T) {
	got := V3Hello()
	want := "Hello, world."
	if got != want {
		t.Errorf("Hello()=%q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	got := Proverb()
	want := "Concurrency is not parallelism."
	if got != want {
		t.Errorf("Proverb()=%q, want %q", got, want)
	}
}
