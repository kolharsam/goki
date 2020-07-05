// NOTE: These tests might be absolute bogus, but if you can let me know of it
// I can change it and learn something in the process! :)

package goki

import "testing"

// Tests
func TestValidateNArgs(t *testing.T) {
	// Test 1
	args := []string{"key", "value"}
	numArgs := ValidateNArgs(2, args)
	if numArgs != nil {
		t.Errorf("ValidateNArgs(2, args) got %d args; want at least %d", len(args), 2)
	}

	// Test 2
	args = []string{"key"}
	numArgs = ValidateNArgs(2, args)
	if numArgs == nil {
		t.Errorf("ValidateNArgs(2, args) got %d args; want at least %d", len(args), 2)
	}

	// Test 3
	args = []string{"key", "value", "some"}
	numArgs = ValidateNArgs(2, args)
	if numArgs == nil {
		t.Errorf("ValidateNArgs(2, args) got %d args; want at least %d", len(args), 2)
	}
}

func TestValidateMinimumNArgs(t *testing.T) {
	// Test 1
	args := []string{"key", "value", "something_else"}
	numArgs := ValidateMinimumNArgs(2, args)
	if numArgs != nil {
		t.Errorf("ValidateMinimumNArgs(2, args) got %d args; want at least %d", len(args), 2)
	}

	// Test 2
	args = []string{"key", "value"}
	numArgs = ValidateMinimumNArgs(3, args)
	if numArgs == nil {
		t.Errorf("ValidateMinimumNArgs(3, args) got %d args. want at least %d", len(args), 3)
	}

	// Test 3
	args = []string{"key", "value"}
	numArgs = ValidateMinimumNArgs(2, args)
	if numArgs != nil {
		t.Errorf("ValidateMinimumNArgs(2, args) got %d args; want at least %d", len(args), 2)
	}
}
