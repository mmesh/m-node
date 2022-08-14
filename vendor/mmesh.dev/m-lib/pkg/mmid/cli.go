package mmid

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const mmCLIIDPrefix string = "__cli"

type MMCLIID string

func NewMMCLIID(accountID, hostID string) MMCLIID {
	mmCLIID := fmt.Sprintf("%s::::%s+%s_%d_%d", accountID, mmCLIIDPrefix, hostID, os.Getpid(), time.Now().Unix())

	return MMCLIID(mmCLIID)
}

func (mmid MMCLIID) String() string {
	return string(mmid)
}

func IsMMCLIID(mmID string) bool {
	s := strings.Split(mmID, ":")
	if len(s) < 5 {
		return false
	}
	return strings.HasPrefix(s[4], mmCLIIDPrefix)
}
