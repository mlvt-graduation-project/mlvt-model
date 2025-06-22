package notify

import (
	"encoding/json"
	"fmt"
	"io"
	"mlvt-api/config"
	"net/http"
	"net/url"
)

// Telegram API structs
type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Text string `json:"text"`
	} `json:"message"`
}

type GetUpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

// GetChatID retrieves recent messages and displays chat IDs
func GetChatID() error {
	cfg := config.Load()

	if cfg.Telegram.BotToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN not set")
	}

	// Get updates from Telegram
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", cfg.Telegram.BotToken)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting updates: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	var updates GetUpdatesResponse
	if err := json.Unmarshal(body, &updates); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	fmt.Println("Recent messages to your bot:")
	fmt.Println("================================")

	if len(updates.Result) == 0 {
		fmt.Println("No messages found. Please send a message to your bot first!")
		fmt.Println("1. Open Telegram")
		fmt.Println("2. Search for your bot")
		fmt.Println("3. Send any message to it")
		fmt.Println("4. Run this command again")
		return nil
	}

	for _, update := range updates.Result {
		fmt.Printf("Chat ID: %d\n", update.Message.Chat.ID)
		fmt.Printf("From: %s (@%s)\n", update.Message.From.FirstName, update.Message.From.Username)
		fmt.Printf("Message: %s\n", update.Message.Text)
		fmt.Println("---")
	}

	// Get the most recent chat ID
	if len(updates.Result) > 0 {
		latestChatID := updates.Result[len(updates.Result)-1].Message.Chat.ID
		fmt.Printf("\n🎯 Use this Chat ID in your .env file:\n")
		fmt.Printf("TELEGRAM_CHAT_ID=%d\n", latestChatID)
	}

	return nil
}

// SendTelegram sends a message to the configured Telegram chat
func SendTelegram(message string) error {
	cfg := config.Load()

	if cfg.Telegram.BotToken == "" || cfg.Telegram.ChatID == "" {
		return fmt.Errorf("telegram configuration not set: missing bot token or chat ID")
	}

	fmt.Println(cfg.Telegram.BotToken)
	fmt.Println(cfg.Telegram.ChatID)

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.Telegram.BotToken)

	data := url.Values{}
	data.Set("chat_id", cfg.Telegram.ChatID)
	data.Set("text", message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("telegram API returned non-200 status: %s", resp.Status)
	}

	return nil
}
