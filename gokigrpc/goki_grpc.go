package gokigrpc

import (
	"errors"

	"golang.org/x/net/context"
)

type Server struct{}

// func (s *Server) RunCommand(ctx context.Context, in *GokiGRPCRequest) (*GokiGRPCResponse, error) {

// }

// Ping is analogous to the Ping service defined for our GRPC server
func (s *Server) Ping(ctx context.Context, in *PingPong) (*PingPong, error) {
	text := in.Text

	if text != "PING" {
		return &PingPong{}, errors.New("inappropriate request")
	}

	return &PingPong{Text: "PONG"}, nil
}
