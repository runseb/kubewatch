package main

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot/slack"

        _ "github.com/runseb/kubewatch/pkg/services"
)

func main() {

	if err := os.Getenv("KUBEBOT_SLACK_TOKEN"); err == "" {
		fmt.Printf("Missing Slack Token. \n")
		return
	}

	slack.Run(os.Getenv("KUBEBOT_SLACK_TOKEN"))
}
