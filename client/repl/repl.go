package repl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kolharsam/goki/client/utils"
	"github.com/kolharsam/goki/common"
)

// RunREPL starts REPL for the client to interact with
func RunREPL(port string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("client:" + port + " >: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if input != "\n" {
			if err = execCommand(input, port); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func execCommand(input string, port string) error {
	command, args, err := parseInput(input)
	if err != nil {
		return err
	}
	request := common.GokiRequest{
		Command: command,
		Args:    args,
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	requestBody := bytes.NewBuffer(jsonBody)
	url := utils.MakeLocalhostURL(port)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return err
	}

	var resp common.GokiResponse
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Result.ToFormat {
		fmt.Fprintln(os.Stdout, "\""+resp.Result.Value+"\"")
	} else {
		fmt.Fprintln(os.Stdout, resp.Result.Value)
	}

	return nil
}

func parseInput(input string) (string, []string, error) {
	cleanInput := utils.CleanseText(input)

	// whitespaces are the delimiters
	inputs := strings.Split(cleanInput, " ")
	if len(inputs) <= 1 {
		return "", []string{}, errors.New("incorrect args passed")
	}

	command := inputs[0]
	args := inputs[1:]

	cleanArgs := []string{}

	for _, arg := range args {
		cleanArgs = append(cleanArgs, utils.CleanseText(arg))
	}

	return command, cleanArgs, nil
}
