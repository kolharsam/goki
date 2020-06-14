package goki

import (
	"errors"
	"strconv"
)

// ValidateNArgs checks if exactly N args have been passed
func ValidateNArgs(numberArgs int, args []string) error {
	currentArgs := len(args)
	if currentArgs != numberArgs {
		return errors.New("Incorrect number of args passed (" + strconv.Itoa(currentArgs) + ") instead of (" + strconv.Itoa(currentArgs) + ")")
	}

	return nil
}

// ValidateMinimumNArgs checks if at least N args have been passed
// perhaps useful for methods with variadic arguments
func ValidateMinimumNArgs(minnum int, args []string) error {
	currentArgs := len(args)
	if currentArgs < minnum {
		return errors.New("Command requires at least " + strconv.Itoa(minnum) + " arguments. You passed (" + strconv.Itoa(currentArgs) + ")")
	}

	return nil
}
