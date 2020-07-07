package cli

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/kolharsam/goki/gokigrpc"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	defer connection.Close()

	client := gokigrpc.NewGokiGRPCClient(conn)

	response, err := client.Ping(context.Background(), &gokigrpc.PingPong{Text: "PING"})
	if err != nil {
		log.Fatalf("Error when calling Ping: %s", err)
	}
	log.Printf("Response from server: %s", response.GetText())
}
