package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func FileParser(filePath string, dst interface{}) error {
	if len(filePath) == 0 {
		return fmt.Errorf("Invalid file path")
	}

	f := viper.New()

	// f.SetConfigType("yaml")
	f.SetConfigFile(filePath)

	if err := f.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return fmt.Errorf("file not found: %v", err)
		} else {
			// Config file was found but another error was produced
			return fmt.Errorf("invalid file: %v", err)
		}
	}

	if err := f.Unmarshal(dst); err != nil {
		return fmt.Errorf("unable to unmarshal file: %v", err)
	}

	return nil
}
