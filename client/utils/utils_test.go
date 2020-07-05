package utils

import "testing"

func TestMakeLocalhostURL(t *testing.T) {
	a := MakeLocalhostURL("9000")
	res := "http://localhost:9000/"
	if a != res {
		t.Errorf("Incorrect url result. Got %s, want %s", a, res)
	}
}

func TestCleanseText(t *testing.T) {
	text := "    set key value          \n"
	a := CleanseText(text)
	res := "set key value"
	if a != res {
		t.Errorf("Could not cleanse text. Got %s, want %s", a, res)
	}
}

func TestContains(t *testing.T) {
	// Test 1
	arr := []string{"key", "keys", "keychain", "keystring"}
	a := Contains(arr, "key")
	if !a {
		t.Error("Incorrect result. Got false. want true")
	}

	// Test 2
	a = Contains(arr, "keycode")
	if a {
		t.Error("Incorrect result. Got true. want false")
	}
}
