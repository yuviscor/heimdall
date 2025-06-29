package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type TelegramNotifier struct {
	ChatID   string
	BotToken string
}

func NewTelegramNotifier(chatId string, botToken string) *TelegramNotifier {
	return &TelegramNotifier{
		ChatID:   chatId,
		BotToken: botToken,
	}
}

func (n *TelegramNotifier) Notify(serviceErr *domain.ServiceError) error {
	sb := strings.Builder{}

	currentTime := time.Now().Format("2006-01-02 15:04:05 UTC")
	sb.WriteString("🚨 <b>SERVICE ALERT</b> 🚨\n")
	sb.WriteString(fmt.Sprintf("⏰ <b>Time:</b> %s\n", currentTime))
	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	sb.WriteString(fmt.Sprintf("🔧 <b>Service Name:</b> <code>%s</code>\n", serviceErr.Name))

	if serviceErr.StatusCode != 0 {
		statusEmoji := "🔴"
		if serviceErr.StatusCode >= 200 && serviceErr.StatusCode < 300 {
			statusEmoji = "🟢"
		} else if serviceErr.StatusCode >= 300 && serviceErr.StatusCode < 400 {
			statusEmoji = "🟡"
		} else if serviceErr.StatusCode >= 400 && serviceErr.StatusCode < 500 {
			statusEmoji = "🟠"
		}
		sb.WriteString(fmt.Sprintf("%s <b>HTTP Status:</b> <code>%d (%s)</code>\n", statusEmoji, serviceErr.StatusCode, getStatusText(serviceErr.StatusCode)))
	}

	if serviceErr.Error != nil {
		sb.WriteString(fmt.Sprintf("❌ <b>Error Details:</b>\n<pre>%v</pre>\n", serviceErr.Error))
	}

	if len(serviceErr.Body) > 0 {
		bodyPreview := string(serviceErr.Body)
		if len(bodyPreview) > 200 {
			bodyPreview = bodyPreview[:200] + "..."
		}
		sb.WriteString(fmt.Sprintf("📄 <b>Response Body:</b>\n<pre>%s</pre>\n", bodyPreview))
	}

	sb.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	sb.WriteString("🔔 <b>Heimdall Monitoring System</b>")

	body := struct {
		ChatID    string `json:"chat_id"`
		Text      string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}{
		ChatID:    n.ChatID,
		Text:      sb.String(),
		ParseMode: "HTML",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.BotToken),
		"application/json",
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error when sending message to telegram: %s", respBody)
	}

	return nil
}
