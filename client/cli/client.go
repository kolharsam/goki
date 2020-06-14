package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goki-client",
	Short:        "cmd to start a client",
	Long:         "Running this command will start a client with REPL to interact with the DB.",
	SilenceUsage: true,
}

var connectCmd = &cobra.Command{
	Use:          "connect",
	Short:        "connect REPL to particular port",
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var port string
		if len(args) == 0 {
			port = "8080"
		} else {
			port = args[0]
		}
		// start the CLIENT REPL here
		ConnectStartREPL(port)
	},
}

// Execute is the main guy leading all of the operation
func Execute() {
	rootCmd.AddCommand(connectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("couldn't execute command", err)
	}
}
