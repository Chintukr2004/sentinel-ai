package alerter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type DiscordAlert struct{}

func Send(message string) error {

	webhook := os.Getenv("DISCORD_WEBHOOK_URL")

	payload := map[string]string{
		"content": message,
	}

	body, _ := json.Marshal(payload)

	_, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))

	return err

}
