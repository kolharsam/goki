package serv

import (
	"errors"
	"fmt"
	"net/http"
)

// StartServerOnPort function returns a new instance of a web server
func StartServerOnPort(port string) {
	serverMux := http.NewServeMux()

	modifiedPort, err := formatPort(port)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	showInitInfo(port)
	mainHandler := http.HandlerFunc(handleServerRequests)
	finalHandler := EnforceJSONPayloadMiddleware(LoggingMiddleware(mainHandler))
	serverMux.Handle("/", finalHandler)
	err = http.ListenAndServe(modifiedPort, serverMux)
	if err != nil {
		fmt.Errorf("there was an error with running the server", err)
	}
}

// showInitInfo is just there to make things fancy...
func showInitInfo(port string) {
	fmt.Printf("Starting goki server on port: %s ...\n", port)
	fmt.Println("Goki-Server v0.0.1")
	fmt.Println("Use the goki-client to connect with this server.")
}

func formatPort(port string) (string, error) {
	if !(len(port) >= 4) && !(len(port) < 6) {
		return "", errors.New("Incorrect port value passed")
	}

	return ":" + port, nil
}

func handleServerRequests(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		// sort of like a health check
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	}

	// TODO: we can probably use something better than a JSON request
	if request.Method == http.MethodPost {
		// var response common.GokiRequest
		// extract the information from the request body - command and args
		// send command and args to the main goki program and get response
		// send the response back to the user
	}
}
