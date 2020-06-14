package goki

import (
	"errors"

	"github.com/kolharsam/goki/common"
)

// var GokiStore

// Execute runs the show for goki and returns
// result of running an operation on goki
func Execute(cmd string, args []string) (common.GokiResponse, error) {
	switch cmd {
	case "set":
		return set(args)
	case "get":
		return get(args)
	default:
		return ({}, errors.New("Command not found"))
	}
}

func set(args []string) (common.GokiResponse, error) {

}

func get(args []string) (common.GokiResponse, error) {

}
