package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func (n *DiscordNotifier) Notify(serviceErr *domain.ServiceError) error {
	sb := strings.Builder{}

	sb.WriteString("🚨 **Service Alert** 🚨\n\n")
	sb.WriteString(fmt.Sprintf("🔧 Service: %s\n", serviceErr.Name))

	if serviceErr.StatusCode != 0 {
		sb.WriteString(fmt.Sprintf("📊 Status Code: %d\n", serviceErr.StatusCode))
	}

	if len(serviceErr.Body) > 0 {
		sb.WriteString(fmt.Sprintf("📄 Response body: %s\n", serviceErr.Body))
	}

	if serviceErr.Error != nil {
		sb.WriteString(fmt.Sprintf("❌ Error: %v\n", serviceErr.Error))
	}

	body := struct {
		Content string `json:"content"`
	}{
		Content: sb.String(),
	}
	bodyBytes, err := json.Marshal(body)
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
