package serv

import (
	"fmt"
	"net/http"
)

// StartServerOnPort function returns a new instance of a web server
func StartServerOnPort(port string) {
	// TODO: validate port
	http.HandleFunc("/", handleServerRequests)
	modifiedPort := attachColon(port)
	showInitInfo(port)

	err := http.ListenAndServe(modifiedPort, nil)
	if err != nil {
		fmt.Errorf("there was an error with running the server", err)
	}
}

// showInitInfo is just there to make things fancy...
func showInitInfo(port string) {
	fmt.Println("Starting goki server on port: ", port, "...")
	fmt.Println("Goki-Server v0.0.1")
	fmt.Println("Use the goki-client to connect with this server.")
}

func attachColon(port string) string {
	return ":" + port
}

func handleServerRequests(writer http.ResponseWriter, request *http.Request) {
	// TODO: handle server requests here
}
