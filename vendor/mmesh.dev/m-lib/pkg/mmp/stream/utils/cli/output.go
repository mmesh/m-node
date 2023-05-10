package cli

import (
	"fmt"

	"mmesh.dev/m-lib/pkg/utils/colors"
)

func Connected() {
	fmt.Printf("%s\n\n", colors.InvertedDarkGreen("Connected"))
}

func Disconnected() {
	fmt.Printf("\n\n%s\n\n", colors.InvertedDarkRedAlt("Disconnected"))
}

/*
func ChatUserLocal(tm int64, user string) string {
	return fmt.Sprintf("\n%s %s%s%s", timestamp(tm), colors.Green("<"), colors.DarkGreen(user), colors.Green(">"))
}

func ChatUserRemote(tm int64, user string) string {
	return fmt.Sprintf("\n%s %s%s%s", timestamp(tm), colors.Yellow("<"), colors.DarkYellow(user), colors.Yellow(">"))
	// return fmt.Sprintf("\n%s %s%s%s", timestamp(tm), colors.Red("<"), colors.DarkRed(user), colors.Red(">"))
}

func timestamp(tm int64) string {
	t := time.Unix(tm, 0)
	hr := t.Hour()
	min := t.Minute()
	sec := t.Second()

	return colors.Black(fmt.Sprintf("%02d:%02d:%02d", hr, min, sec))
}
*/
