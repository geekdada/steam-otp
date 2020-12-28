package main

import (
	"fmt"
	"github.com/geekdada/steam-otp/v2/otp"
	"github.com/jessevdk/go-flags"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func main() {
	var options struct {
		Secret string `short:"s" long:"secret" description:"The shared secret you get from SteamDesktopAuthenticator"`
	}

	_, err := flags.Parse(&options)

	if err != nil {
		switch err.(type) {
		case *flags.Error:
			fe, _ := err.(*flags.Error)
			if fe.Type == flags.ErrHelp {
				os.Exit(0)
			}
			fmt.Println(err)
			os.Exit(1)
		default:
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if options.Secret != "" {
		code, err := otp.GenerateAuthCode(strings.TrimSpace(options.Secret), 0)

		if err != nil {
			fmt.Printf("Failed to generate code %v\n", err)
		}

		fmt.Println(code)

		return
	}

	prompt := promptui.Prompt{
		Label: "Input the shared secret you get from SteamDesktopAuthenticator",
		Mask:  '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Invalid input %v\n", err)
		return
	}

	code, err := otp.GenerateAuthCode(strings.TrimSpace(result), 0)

	if err != nil {
		fmt.Printf("Failed to generate code %v\n", err)
	}

	fmt.Println(code)
}
