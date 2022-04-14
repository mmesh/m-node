package mmid

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const mmCLIIDPrefix string = "__cli"

type MMCLIID string

func NewMMCLIID(hostID string) MMCLIID {
	mmCLIID := fmt.Sprintf("%s:%s:%d:%d", mmCLIIDPrefix, hostID, os.Getpid(), time.Now().Unix())

	return MMCLIID(mmCLIID)
}

func (mmid MMCLIID) String() string {
	return string(mmid)
}

func IsMMCLIID(mmID string) bool {
	return strings.Split(mmID, ":")[0] == mmCLIIDPrefix
}
