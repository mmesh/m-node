package metrics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const netDevFile string = "/proc/net/dev"
const rmsIface string = "eth0"

type Iface string

type networkStat struct {
	Iface        string `json:"iface"`
	RxBytes      uint64 `json:"rxBytes"`
	RxPackets    uint64 `json:"rxPackets"`
	RxErrors     uint64 `json:"rxErrors"`
	RxDrops      uint64 `json:"rxDrops"`
	RxFifo       uint64 `json:"rxFifo"`
	RxFrame      uint64 `json:"rxFrame"`
	RxCompressed uint64 `json:"rxCompressed"`
	RxMulticast  uint64 `json:"rxMulticast"`
	TxBytes      uint64 `json:"txBytes"`
	TxPackets    uint64 `json:"txPackets"`
	TxErrors     uint64 `json:"txErrors"`
	TxDrops      uint64 `json:"txDrops"`
	TxFifo       uint64 `json:"txFifo"`
	TxColls      uint64 `json:"txCollisions"`
	TxCarrier    uint64 `json:"txCarrier"`
	TxCompressed uint64 `json:"txCompressed"`
}

type networkStatMap struct {
	prevStats map[Iface]*networkStat
	stats     map[Iface]*networkStat
	timestamp int64
	sync.RWMutex
}

var nsMap *networkStatMap

func newNetworkStatMap() *networkStatMap {
	return &networkStatMap{
		prevStats: make(map[Iface]*networkStat),
		stats:     make(map[Iface]*networkStat),
	}
}

func ReadNetDevStats() error {
	file, err := os.Open(netDevFile)
	if err != nil {
		return err
	}
	defer file.Close()

	s := &networkStat{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		kv := strings.SplitN(scanner.Text(), ":", 2)
		if len(kv) != 2 {
			continue
		}

		fields := strings.Fields(kv[1])
		if len(fields) < 16 {
			continue
		}

		s.Iface = strings.TrimSpace(kv[0])
		if s.Iface == "lo" {
			continue
		}

		var err error

		s.RxBytes, err = strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxBytes of %s", s.Iface)
		}
		s.RxPackets, err = strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxPackets of %s", s.Iface)
		}
		s.RxErrors, err = strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxErrors of %s", s.Iface)
		}
		s.RxDrops, err = strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxDrops of %s", s.Iface)
		}
		s.RxFifo, err = strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxFifo of %s", s.Iface)
		}
		s.RxFrame, err = strconv.ParseUint(fields[5], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxFrame of %s", s.Iface)
		}
		s.RxCompressed, err = strconv.ParseUint(fields[6], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxCompressed of %s", s.Iface)
		}
		s.RxMulticast, err = strconv.ParseUint(fields[7], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse rxMulticast of %s", s.Iface)
		}

		s.TxBytes, err = strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txBytes of %s", s.Iface)
		}
		s.TxPackets, err = strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txPackets of %s", s.Iface)
		}
		s.TxErrors, err = strconv.ParseUint(fields[10], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txErrors of %s", s.Iface)
		}
		s.TxDrops, err = strconv.ParseUint(fields[11], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txDrops of %s", s.Iface)
		}
		s.TxFifo, err = strconv.ParseUint(fields[12], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txFifo of %s", s.Iface)
		}
		s.TxColls, err = strconv.ParseUint(fields[13], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txCollisions of %s", s.Iface)
		}
		s.TxCarrier, err = strconv.ParseUint(fields[14], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txCarrier of %s", s.Iface)
		}
		s.TxCompressed, err = strconv.ParseUint(fields[15], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse txCompressed of %s", s.Iface)
		}

		if nsMap == nil {
			nsMap = newNetworkStatMap()
		}
		nsMap.update(s)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan error for /proc/net/dev: %s", err)
	}

	return nil
}

func (n *networkStatMap) update(s *networkStat) {
	n.Lock()
	defer n.Unlock()

	// if n.timestamp > 0 {
	// 	tm := time.Now().Unix()
	// 	delta := n.timestamp - tm
	// }

	iface := Iface(s.Iface)

	if p, ok := n.prevStats[iface]; ok {
		n.stats[iface] = &networkStat{
			RxBytes:      s.RxBytes - p.RxBytes,
			RxPackets:    s.RxPackets - p.RxPackets,
			RxErrors:     s.RxErrors - p.RxErrors,
			RxDrops:      s.RxDrops - p.RxDrops,
			RxFifo:       s.RxFifo - p.RxFifo,
			RxFrame:      s.RxFrame - p.RxFrame,
			RxCompressed: s.RxCompressed - p.RxCompressed,
			RxMulticast:  s.RxMulticast - p.RxMulticast,
			TxBytes:      s.TxBytes - p.TxBytes,
			TxPackets:    s.TxPackets - p.TxPackets,
			TxErrors:     s.TxErrors - p.TxErrors,
			TxDrops:      s.TxDrops - p.TxDrops,
			TxFifo:       s.TxFifo - p.TxFifo,
			TxColls:      s.TxColls - p.TxColls,
			TxCarrier:    s.TxCarrier - p.TxCarrier,
			TxCompressed: s.TxCompressed - p.TxCompressed,
		}

		if s.Iface == rmsIface {
			setNetDevStats(
				n.stats[iface].TxBytes,
				n.stats[iface].RxBytes,
				n.stats[iface].TxPackets,
				n.stats[iface].RxPackets,
			)
		}
	}

	n.prevStats[iface] = s

	n.timestamp = time.Now().Unix()
}
