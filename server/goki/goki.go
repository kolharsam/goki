package goki

import (
	"errors"
	"time"

	"github.com/kolharsam/goki/common"
)

// Execute runs the show for goki and returns
// result of running an operation on goki
func Execute(cmd string, args []string) (common.GokiResponse, error) {
	switch cmd {
	case "set":
		return set(args)
	case "get":
		return get(args)
	default:
		return common.GokiResponse{}, errors.New("Command not found")
	}
}

func set(args []string) (common.GokiResponse, error) {
	err := ValidateNArgs(2, args)
	if err != nil {
		return common.GokiResponse{}, err
	}

	key := args[0]
	value := args[1]

	result := SetToStorage(key, value)

	return common.GokiResponse{
		Result:    result,
		TimeStamp: time.Now(),
	}, nil
}

func get(args []string) (common.GokiResponse, error) {
	err := ValidateNArgs(1, args)
	if err != nil {
		return common.GokiResponse{}, err
	}

	key := args[0]
	value := GetValue(key)

	return common.GokiResponse{
		Result:    value,
		TimeStamp: time.Now(),
	}, nil
}
