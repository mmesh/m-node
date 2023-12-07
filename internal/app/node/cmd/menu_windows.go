//go:build windows
// +build windows

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"mmesh.dev/m-lib/pkg/logging"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/start"
)

// serviceStartCmd represents the service-start command
var serviceStartCmd = &cobra.Command{
	Use:   "service-start",
	Short: "Start Windows service",
	Long:  `Start Windows service.`,
	Run: func(cmd *cobra.Command, args []string) {
		xlog.Logger().SetANSIColor(false)
		xlog.Logger().SetLogFile(logFile())
		xlog.Logger().SetWindowsLogger(&xlog.EventLogOptions{
			Level:  logging.LogLevel,
			Source: version.NODE_NAME,
		})

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

		xlog.Logger().SetStdLogger()

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

		xlog.Logger().SetStdLogger()

		start.ServiceUninstall()
	},
}

func init() {
	rootCmd.AddCommand(serviceStartCmd)
	rootCmd.AddCommand(serviceInstallCmd)
	rootCmd.AddCommand(serviceUninstallCmd)
}
