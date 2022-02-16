package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	godotenv.Load(".env")

	token := os.Getenv("AUTH_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	app_token := os.Getenv("APP_TOKEN")

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(app_token))
	// Attachment displayed from the bot
	attachment := slack.Attachment{
		Pretext: "New message from bot",
		Text:    "Hello from vbot",
		Color:   "#ba2507",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().Format("2006-01-02 15:04:05 MST"),
			},
		},
	}

	// Connecting to socketmode
	socket := socketmode.New(
		client,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode", log.Lshortfile|log.LstdFlags)),
	)
	socket.Run()

	// Sends the message
	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
}
