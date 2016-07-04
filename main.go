package main

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot/slack"

        _ "github.com/runseb/kubebot/pkg/services"
)

const (
	SlackToken        string = "KUBEBOT_SLACK_TOKEN"
)

func main() {

	if err := os.Getenv(SlackToken); err == "" {
		fmt.Printf("Missing Slack Token. \n")
		return
	}

	slack.Run(os.Getenv(SlackToken))
}
