package config

import (
	"testing"

	"github.com/MowlCoder/heimdall/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Enablement_Notifiers(t *testing.T) {
	cfg := &Config{
		Services: []domain.Service{
			{Name: "My Personal Website", URL: "https://example.com"},
		},
		Notifiers: NotifiersConfig{
			Telegram: &TelegramNotifierConfig{
				ChatID:   "-123456789",
				BotToken: "secret-telegram-bot-token",
				Enabled:  true,
			},
			Slack: &SlackNotifierConfig{
				ChatID:   "RandomTest",
				BotToken: "xoxosecret-slack-bot-token",
				Enabled:  true,
			},
			Discord: &DiscordNotifierConfig{
				Webhook: "random",
				Enabled: true,
			},
			Webhook: &WebhookNotifierConfig{
				Webhook: "randomWebhook",
				Enabled: true,
			},
		},
	}

	checkTelegram := cfg.IsTelegramEnabled()
	assert.True(t, checkTelegram)
	checkSlack := cfg.IsSlackEnabled()
	assert.True(t, checkSlack)
	checkDiscord := cfg.IsDiscordEnabled()
	assert.True(t, checkDiscord)
	checkWebhook := cfg.IsWebhookEnabled()
	assert.True(t, checkWebhook)
}
