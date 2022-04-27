package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-node/internal/app/node/start"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the agent",
	Long:  `Start the agent.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ConsoleInit(); err != nil {
			log.Fatal(err)
		}

		start.Main()
	},
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Show the ` + version.NAME + ` client version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Agent Info: " + version.NODE_NAME + " " + version.GetVersion() + "\n")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
}
