package goki

import (
	"errors"

	"github.com/kolharsam/goki/common"
)

var commands = map[string]func([]string) (common.GokiResponse, error){
	"set":    Set,
	"get":    Get,
	"exists": Exists,
	"del":    Delete,
	"expire": Expire,
	"ttl":    TTL,
	"incr":   Increment,
	"decr":   Decrement,
}

// Execute runs the show for goki and returns
// result of running an operation on goki
func Execute(cmd string, args []string) (common.GokiResponse, error) {
	if fn, ok := commands[cmd]; ok != false {
		return fn(args)
	}
	return common.GokiResponse{}, errors.New("Command not found")
}
