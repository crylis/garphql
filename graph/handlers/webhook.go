package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const discordWebhookURL = "https://discord.com/api/webhooks/1197086686576906281/KWHJ7bPUOKzlDQgz9nkIR57XVx8SaytJ5EP7CeqUwhHA05uwatF5I1L-fp77fTdqyQ28"

type DiscordMessage struct {
	Content string `json:"content"`
}

func SendDiscordMessage(message string) error {
	webhookMessage := DiscordMessage{Content: message}

	payload, err := json.Marshal(webhookMessage)
	if err != nil {
		return err
	}

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Discord message, status code: %d", resp.StatusCode)
	}

	return nil
}
