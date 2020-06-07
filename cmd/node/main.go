package main

import (
	"fmt"

	"mmesh.dev/mmesh/internal/app/node/cmd"
	"mmesh.dev/mmesh/internal/pkg/version"
	"x6a.dev/pkg/colors"
)

func main() {
	fmt.Println(colors.White(version.AGENT_NAME + " " + version.GetVersion() + "\n"))

	cmd.Execute()
}
