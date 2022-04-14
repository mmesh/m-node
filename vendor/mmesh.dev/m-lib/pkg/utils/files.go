package utils

import (
	"os"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		// filePath exists
		//logging.Debugf("File %s found", filePath)
		return true
	} else if os.IsNotExist(err) {
		// filePath does *not* exist
		//logging.Debug(err)
		return false
	} else {
		// filePath may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		//logging.Error(err)
		return false
	}
}
