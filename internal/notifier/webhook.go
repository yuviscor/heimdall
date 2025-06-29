package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type WebhookNotifier struct {
	Webhook string
}

func NewWebhookNotifier(webhook string) *WebhookNotifier {
	return &WebhookNotifier{
		Webhook: webhook,
	}
}

func (n *WebhookNotifier) Notify(serviceErr *domain.ServiceError) error {
	bodyBytes, err := json.Marshal(serviceErr)
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
		return fmt.Errorf("error when sending message to custom webhook: %s", respBody)
	}

	return nil
}
