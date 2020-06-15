package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kolharsam/goki/client/repl"
	"github.com/kolharsam/goki/client/utils"
)

// ConnectStartREPL starts a client REPL upon verifying that a valid server is running
func ConnectStartREPL(port string) {
	err := ping(port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	showInitInfo(port)
	repl.RunREPL(port)
}

// ping checks whether there's a server
func ping(port string) error {
	url := utils.MakeLocalhostURL(port)
	response, err := http.Get(url)
	if err != nil {
		return errors.New("Unable to connect to server at :" + port)
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("Unable to connect to server at :" + port)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Can't seem to read the server response")
	}

	fmt.Println(string(body))
	return nil
}

func showInitInfo(port string) {
	fmt.Println("Connecting to server at port: " + port)
	fmt.Println("Goki-Client v0.0.1\n\nStarting REPL client...")
}
