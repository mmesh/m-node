package mmp

import (
	"fmt"

	"x6a.dev/pkg/colors"
)

/*
func WelcomeToTheNetwork() {
	text := "Welcome to mmesh network"
	prePad := strings.Repeat("─", 3)
	postPad := strings.Repeat("─", 70-len(text)-1)
	// pad := strings.Repeat("█", 60-len(text)-1)
	fmt.Printf("\n%s %s %s\n", colors.DarkCyan(prePad), colors.Black(text), colors.DarkCyan(postPad))
}

func endOfTransmission() {
	text := "End-of-Transmission"
	prePad := strings.Repeat("─", 3)
	postPad := strings.Repeat("─", 70-len(text)-1)
	fmt.Printf("\n%s %s %s\n", colors.Black(prePad), colors.Black(text), colors.Black(postPad))
}
*/

func Connected() {
	fmt.Printf("%s\n\n", colors.InvertedGreen("Connected"))
}

func Disconnected() {
	fmt.Printf("\n%s\n\n", colors.InvertedBlack("Disconnected"))
}
