package goki

import (
	"time"

	"github.com/kolharsam/goki/common"
)

// Set is used to set a key into the data store
func Set(args []string) (common.GokiResponse, error) {
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

// Get returns the value stored in a particular key, if key is not found
// it returns an error or nil
func Get(args []string) (common.GokiResponse, error) {
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

// Exists returns true if a key exists in the database
func Exists(args []string) (common.GokiResponse, error) {
	err := ValidateNArgs(1, args)
	if err != nil {
		return common.GokiResponse{}, err
	}

	key := args[0]
	value := CheckIfExists(key)

	return common.GokiResponse{
		Result:    value,
		TimeStamp: time.Now(),
	}, nil
}

// Delete returns OK upon deleting a key if it is present in the db
func Delete(args []string) (common.GokiResponse, error) {
	err := ValidateNArgs(1, args)
	if err != nil {
		return common.GokiResponse{}, err
	}

	key := args[0]
	value := DeleteKey(key)

	return common.GokiResponse{
		Result:    value,
		TimeStamp: time.Now(),
	}, nil
}

// Expire returns OK upon deleting a key if it is present in the db
func Expire(args []string) (common.GokiResponse, error) {
	// TODO: This can be used in conjunction with set
	// eg. set hello ex 100, so we should add support for such a possibility
	err := ValidateNArgs(2, args)
	if err != nil {
		return common.GokiResponse{}, err
	}

	key := args[0]
	expTime := args[1]
	ExpireKey(key, expTime)

	return common.GokiResponse{}, nil
}
