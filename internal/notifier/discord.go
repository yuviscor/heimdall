package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type DiscordNotifier struct {
	Webhook string
}

func NewDiscordNotifier(webhook string) *DiscordNotifier {
	return &DiscordNotifier{
		Webhook: webhook,
	}
}

type discordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Color       int                 `json:"color"`
	Fields      []discordEmbedField `json:"fields"`
	Timestamp   string              `json:"timestamp"`
	Footer      discordEmbedFooter  `json:"footer"`
}

type discordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type discordEmbedFooter struct {
	Text string `json:"text"`
}

type discordWebhookPayload struct {
	Content string         `json:"content"`
	Embeds  []discordEmbed `json:"embeds"`
}

func (n *DiscordNotifier) Notify(serviceErr *domain.ServiceError) error {
	currentTime := time.Now()

	color := 0xFF0000
	if serviceErr.StatusCode >= 200 && serviceErr.StatusCode < 300 {
		color = 0x00FF00
	} else if serviceErr.StatusCode >= 300 && serviceErr.StatusCode < 400 {
		color = 0xFFFF00
	} else if serviceErr.StatusCode >= 400 && serviceErr.StatusCode < 500 {
		color = 0xFFA500
	}

	fields := []discordEmbedField{
		{
			Name:   "🔧 Service Name",
			Value:  fmt.Sprintf("`%s`", serviceErr.Name),
			Inline: true,
		},
	}

	if serviceErr.StatusCode != 0 {
		statusText := getStatusText(serviceErr.StatusCode)
		fields = append(fields, discordEmbedField{
			Name:   "📊 HTTP Status",
			Value:  fmt.Sprintf("`%d (%s)`", serviceErr.StatusCode, statusText),
			Inline: true,
		})
	}

	fields = append(fields, discordEmbedField{
		Name:   "⏰ Timestamp",
		Value:  fmt.Sprintf("<t:%d:F>", currentTime.Unix()),
		Inline: true,
	})

	if serviceErr.Error != nil {
		errorDetails := fmt.Sprintf("%v", serviceErr.Error)
		if len(errorDetails) > 1024 {
			errorDetails = errorDetails[:1021] + "..."
		}
		fields = append(fields, discordEmbedField{
			Name:   "❌ Error Details",
			Value:  fmt.Sprintf("```%s```", errorDetails),
			Inline: false,
		})
	}

	if len(serviceErr.Body) > 0 {
		bodyPreview := string(serviceErr.Body)
		if len(bodyPreview) > 1024 {
			bodyPreview = bodyPreview[:1021] + "..."
		}
		fields = append(fields, discordEmbedField{
			Name:   "📄 Response Body",
			Value:  fmt.Sprintf("```%s```", bodyPreview),
			Inline: false,
		})
	}

	embed := discordEmbed{
		Title:       "🚨 SERVICE ALERT 🚨",
		Description: "A service monitoring alert has been triggered",
		Color:       color,
		Fields:      fields,
		Timestamp:   currentTime.Format(time.RFC3339),
		Footer: discordEmbedFooter{
			Text: "Heimdall Monitoring System",
		},
	}

	payload := discordWebhookPayload{
		Content: "🔔 **Service Alert Detected**",
		Embeds:  []discordEmbed{embed},
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(n.Webhook, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error when sending message to discord: %s", respBody)
	}

	return nil
}
