package main

import (
	"fmt"
	"os"

	"github.com/go-chat-bot/bot/slack"
)

const (
	slackTokenLabel        string = "KUBEBOT_SLACK_TOKEN"
)

var (
	kb *Kubebot
)

func main() {

	if err := os.Getenv(slackTokenLabel); err == "" {
		fmt.Printf("Missing Slack Token. \n")
		return
	}

	kb = &Kubebot{
		token:    os.Getenv(slackTokenLabel),
	}

	slack.Run(kb.token)
}
