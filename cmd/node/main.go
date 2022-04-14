package main

import (
	"fmt"
	// "log"

	"mmesh.dev/m-lib/pkg/version"
	"mmesh.dev/m-node/internal/app/node/cmd"
)

func main() {
	// if err := cmd.ConsoleInit(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%s %s ", version.NODE_NAME, version.GetVersion())

	cmd.Execute()
}
