//go:build windows
// +build windows

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"mmesh.dev/m-node/internal/app/node/start"
)

// serviceStartCmd represents the service-start command
var serviceStartCmd = &cobra.Command{
	Use:   "service-start",
	Short: "Start Windows service",
	Long:  `Start Windows service.`,
	Run: func(cmd *cobra.Command, args []string) {
		start.ServiceStart()
	},
}

// serviceInstallCmd represents the service-install command
var serviceInstallCmd = &cobra.Command{
	Use:   "service-install",
	Short: "Install Windows service",
	Long:  `Install Windows service.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ConsoleInit(); err != nil {
			log.Fatal(err)
		}

		start.ServiceInstall()
	},
}

// serviceUninstallCmd represents the service-uninstall command
var serviceUninstallCmd = &cobra.Command{
	Use:   "service-uninstall",
	Short: "Uninstall Windows service",
	Long:  `Uninstall Windows service.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ConsoleInit(); err != nil {
			log.Fatal(err)
		}

		start.ServiceUninstall()
	},
}

func init() {
	rootCmd.AddCommand(serviceStartCmd)
	rootCmd.AddCommand(serviceInstallCmd)
	rootCmd.AddCommand(serviceUninstallCmd)
}
