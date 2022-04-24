package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/utils/msg"
	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-node/internal/app/node/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   version.NODE_SHORTNAME,
	Short: version.NODE_SHORTNAME + " is the agent of the " + version.PLATFORM_NAME,
	Long: version.NODE_SHORTNAME + ` is the agent of the ` +
		version.PLATFORM_NAME + `.

Find support and more information:

  Project Website:     ` + version.MMESH_URL + `
  Documentation:       ` + version.MMESH_DOC_URL + `
  Join Us on Discord:  ` + version.DISCORD_URL,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		msg.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", defaultConfigFile(), "configuration file")
	// rootCmd.PersistentFlags().BoolP("insecure", "", false, "insecure mode (TLS disabled)")
	// rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode")
	// rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "table", "output format (table, json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	if len(cfgFile) > 0 {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.SetEnvPrefix("mm") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv() // read in environment variables that match

	// viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))
	// viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		msg.Errorf("Unable to read configuration file: %s\n", err)
		// os.Exit(1)
	} else {
		// msg.Debugf("Using configuration file: %v\n", viper.ConfigFileUsed())
	}

	config.Init()
}
