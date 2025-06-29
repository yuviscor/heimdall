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

type SlackNotifier struct {
	ChatID   string
	BotToken string
}

func NewSlackNotifier(chatId string, botToken string) *SlackNotifier {
	return &SlackNotifier{
		ChatID:   chatId,
		BotToken: botToken,
	}
}

type slackSendMessageResponse struct {
	OK bool `json:"ok"`
}

func (n *SlackNotifier) Notify(serviceErr *domain.ServiceError) error {
	sb := strings.Builder{}

	currentTime := time.Now().Format("2006-01-02 15:04:05 UTC")
	sb.WriteString("🚨 *SERVICE ALERT* 🚨\n")
	sb.WriteString(fmt.Sprintf("⏰ *Time:* %s\n", currentTime))
	sb.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	sb.WriteString(fmt.Sprintf("🔧 *Service Name:* `%s`\n", serviceErr.Name))

	if serviceErr.StatusCode != 0 {
		statusEmoji := "🔴"
		if serviceErr.StatusCode >= 200 && serviceErr.StatusCode < 300 {
			statusEmoji = "🟢"
		} else if serviceErr.StatusCode >= 300 && serviceErr.StatusCode < 400 {
			statusEmoji = "🟡"
		} else if serviceErr.StatusCode >= 400 && serviceErr.StatusCode < 500 {
			statusEmoji = "🟠"
		}
		sb.WriteString(fmt.Sprintf("%s *HTTP Status:* `%d` (%s)\n", statusEmoji, serviceErr.StatusCode, getStatusText(serviceErr.StatusCode)))
	}

	if serviceErr.Error != nil {
		sb.WriteString(fmt.Sprintf("❌ *Error Details:*\n```%v```\n", serviceErr.Error))
	}

	if len(serviceErr.Body) > 0 {
		bodyPreview := string(serviceErr.Body)
		if len(bodyPreview) > 200 {
			bodyPreview = bodyPreview[:200] + "..."
		}
		sb.WriteString(fmt.Sprintf("📄 *Response Body:*\n```%s```\n", bodyPreview))
	}

	sb.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	sb.WriteString("🔔 *Heimdall Monitoring System*")

	body := struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
	}{
		Channel: n.ChatID,
		Text:    sb.String(),
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", n.BotToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("error when sending message to slack: %s", respBody)
	}

	var parsedResponse slackSendMessageResponse
	if err := json.Unmarshal(respBody, &parsedResponse); err != nil {
		return err
	}

	if !parsedResponse.OK {
		return fmt.Errorf("error when sending message to slack: %s", respBody)
	}

	return nil
}
