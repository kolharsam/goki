package serv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kolharsam/goki/common"
	"github.com/kolharsam/goki/server/goki"
	"github.com/kolharsam/goki/server/logger"
)

// StartServerOnPort function returns a new instance of a web server
func StartServerOnPort(port string) {
	serverMux := http.NewServeMux()

	modifiedPort, err := formatPort(port)
	if err != nil {
		fmt.Println(err.Error())
	}

	showInitInfo(port)
	mainHandler := http.HandlerFunc(handleServerRequests)
	finalHandler := EnforceJSONPayloadMiddleware(LoggingMiddleware(mainHandler))
	serverMux.Handle("/", finalHandler)

	// start the server
	err = http.ListenAndServe(modifiedPort, serverMux)
	if err != nil {
		logger.LogErr(err)
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
		writer.Write([]byte("PONG"))
	}

	// NOTE: we can probably use something better than a JSON request
	if request.Method == http.MethodPost {
		var req common.GokiRequest
		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := goki.Execute(req.Command, req.Args)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	}
}
