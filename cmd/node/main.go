package main

import (
	"fmt"

	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-node/internal/app/node/cmd"
	"x6a.dev/pkg/colors"
)

func main() {
	fmt.Println(colors.White(version.NODE_NAME + " " + version.GetVersion() + "\n"))

	cmd.Execute()
}
