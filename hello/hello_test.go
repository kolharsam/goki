package hello

import "testing"

// HelloWorldTest is a test
func HelloWorldTest(t *testing.T) {
	want := "Hello, world!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
