package output

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func AppHeader(accountID string, output bool) string {
	var account string
	if len(accountID) == 0 {
		accountID = viper.GetString("account.id")
		if len(accountID) > 0 {
			accountID = colors.DarkWhite(fit(accountID, 28))
			account = fmt.Sprintf(" Account ID: %s", accountID)
		}
	} else {
		accountID = colors.DarkWhite(fit(accountID, 28))
		account = fmt.Sprintf(" Account ID: %s", accountID)
	}

	var user string
	userEmail := viper.GetString("user.email")
	if len(userEmail) > 0 {
		userEmail = colors.DarkWhite(fit(userEmail, 28))
		user = fmt.Sprintf(" User:       %s", userEmail)
	}

	endpoint := viper.GetString("controller.authserver")
	if len(endpoint) == 0 {
		endpoint = colors.Black("[not-configured]")
	} else {
		if strings.HasSuffix(endpoint, "mmesh.dev") ||
			strings.HasSuffix(endpoint, "mmesh.network") {
			endpoint = fmt.Sprintf(
				"%s.%s.%s",
				strings.Split(endpoint, ".")[1],
				strings.Split(endpoint, ".")[2],
				strings.Split(endpoint, ".")[3],
			)
		}
		endpoint = colors.DarkWhite(fit(endpoint, 28))
	}

	return logo(account, user, endpoint, output)
}

func logo(account, user, endpoint string, output bool) string {
	if output {
		l1 := fmt.Sprintf(colors.DarkWhite("  ■   ▄  ▄▄ ▄▄ ▄▄ ▄▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ")+colors.Green("│")+"%s\n", account)
		l2 := fmt.Sprintf(colors.DarkWhite("■  ██    █ ▄ █ █ ▄ █ █■   ▀  ▄ █▄▄█ ")+colors.Green("│")+"%s\n", user)
		l3 := fmt.Sprintf(colors.DarkWhite("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ")+colors.Green("│")+" Controller: %s\n", endpoint)
		return fmt.Sprintf("%s%s%s\n", l1, l2, l3)
	}

	fmt.Printf("  ■   ▄  ▄▄ ▄▄ ▄▄ ▄▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ "+colors.Green("│")+"%s\n", account)
	fmt.Printf("■  ██    █ ▄ █ █ ▄ █ █■   ▀  ▄ █▄▄█ "+colors.Green("│")+"%s\n", user)
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ "+colors.Green("│")+" Controller: %s\n", endpoint)
	fmt.Println()

	return ""
}

/*
func Logo0() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄    " + colors.Blue("│") + " \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■▀ ▀  ▄ █■■▄ " + colors.Blue("│") + " \n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ " + colors.Blue("│") + " Discover more at " + colors.DarkWhite("https://mmesh.io") + "\n")
}

func Logo1() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ │                  \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ │ Discover more at:\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ │ https://mmesh.io \n")
}

func Logo2() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ┌┐                 \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ ││ https://mmesh.io\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ └┘                 \n")
}

func Logo3() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄ ┌┐┌───┐                  \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█ │││ ■ │ Discover more at:\n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ └┘└───┘ https://mmesh.io \n")
}

func Logo4() {
	fmt.Printf("  ■   ▄  ▄   ▄ ▄   ▄ ▄▄▄▄ ▄▄▄▄ ▄  ▄   ┌┐┌───┐ \n")
	fmt.Printf("■  ██    █▀▄▀█ █▀▄▀█ █■■  ▀■■▄ █■■█   │││ ■ │ \n")
	fmt.Printf("  ▀   ■  ▀ ▀ ▀ ▀ ▀ ▀ ▀▀▀▀ ▀▀▀▀ ▀  ▀ ▀ └┘└───┘ \n")
}

func logo5() {
	fmt.Print("██  ██ ██  ██ ██████ ██████ ██  ██")
	fmt.Print("██████ ██████ ████   ██████ ██████")
	fmt.Print("██  ██ ██  ██ ██████ ██████ ██  ██")

	fmt.Print("┌┐  ┌┐┌┐  ┌┐┌▄▄┌▄▄ ┌┐  ┌┐")
	fmt.Print("││┌┐││││┌┐│││■ └■■┐││■■││")
	fmt.Print("└┘└┘└┘└┘└┘└┘└▀▀ ▀▀┘└┘  └┘")

	fmt.Print("╔╗┌┐  ┌┐┌┐  ┌┐╔═┌── ┌┐  ┌┐")
	fmt.Print("║║││┌┐││││┌┐││╠═└──┐│────│")
	fmt.Print("╚╝┘└┘└┘└┘└┘└┘└╚═ ──┘└┘  └┘")
}
*/
