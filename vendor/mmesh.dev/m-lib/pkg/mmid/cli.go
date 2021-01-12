package mmid

import (
	"fmt"
	"os"
	"time"
)

type MMCLIID string

func NewMMCLIID(hostID string) MMCLIID {
	mmCLIID := fmt.Sprintf("%s:%d:%d", hostID, os.Getpid(), time.Now().Unix())

	return MMCLIID(mmCLIID)
}

func (mmid MMCLIID) String() string {
	return string(mmid)
}
