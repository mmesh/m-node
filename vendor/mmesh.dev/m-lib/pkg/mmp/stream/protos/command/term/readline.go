package term

import (
	"bufio"
	"os"
	"strings"

	"mmesh.dev/m-lib/pkg/utils/msg"
)

func ReadLine() []byte {
	reader := bufio.NewReader(os.Stdin)
	// line, err := reader.ReadBytes('\n')
	line, err := reader.ReadString('\n')
	if err != nil {
		msg.Alertf("Unable to read from stdin: %v", err)
		return nil
	}

	// fmt.Println(line)

	return []byte(strings.TrimSpace(line))
}
