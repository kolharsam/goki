package serv

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goki-server",
	Short:        "cmd to start server",
	Long:         "Running this command will start a server for you.",
	SilenceUsage: true,
}

var startCmd = &cobra.Command{
	Use:          "start",
	Short:        "start a web server at a particular port",
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var port string
		if len(args) == 0 {
			port = "8080"
		} else {
			port = args[0]
		}
		// start the server here
		StartServerOnPort(port)
	},
}

// Execute is the main guy leading all of the operation
func Execute() {
	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("couldn't execute command", err)
	}
}
