package main

import (
	"fmt"

	"github.com/Goryudyuma/GithHubActionsTest/main/actiontype"
	slack5min "github.com/Goryudyuma/GithHubActionsTest/main/slack/5min"
)

func readArgument() actiontype.Argument {
	return actiontype.Argument{}
}

func main() {
	args := readArgument()

	slack5min.Run(args.GetSlackWebhookUrl5min())

	fmt.Println("Hello golang!")
}
