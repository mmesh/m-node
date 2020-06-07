package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"mmesh.dev/mmesh/internal/app/node/start"
	"mmesh.dev/mmesh/internal/pkg/version"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the agent",
	Long:  `Start the agent.`,
	Run: func(cmd *cobra.Command, args []string) {
		start.Main()
	},
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Show the ` + version.NAME + ` client version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Agent Info: " + version.AGENT_SHORTNAME + " " + version.GetVersion() + "\n")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
}
