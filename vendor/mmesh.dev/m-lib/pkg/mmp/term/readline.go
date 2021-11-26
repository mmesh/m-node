package term

import (
	"bufio"
	"os"
	"strings"

	"x6a.dev/pkg/msg"
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
