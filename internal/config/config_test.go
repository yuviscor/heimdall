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
		Notifiers: struct {
			Telegram *struct {
				ChatID   string "json:\"chatId\""
				BotToken string "json:\"botToken\""
				Enabled  bool   "json:\"enabled\""
			}
			Discord *struct {
				Webhook string "json:\"webhook\""
				Enabled bool   "json:\"enabled\""
			}
			Slack *struct {
				ChatID   string "json:\"chatId\""
				BotToken string "json:\"botToken\""
				Enabled  bool   "json:\"enabled\""
			}
			Webhook *struct {
				Webhook string "json:\"webhook\""
				Enabled bool   "json:\"enabled\""
			}
		}{

			Telegram: &struct {
				ChatID   string "json:\"chatId\""
				BotToken string "json:\"botToken\""
				Enabled  bool   "json:\"enabled\""
			}{

				ChatID:   "-123456789",
				BotToken: "secret-telegram-bot-token",
				Enabled:  true,
			},
			Slack: &struct {
				ChatID   string "json:\"chatId\""
				BotToken string "json:\"botToken\""
				Enabled  bool   "json:\"enabled\""
			}{
				ChatID:   "RandomTest",
				BotToken: "xoxosecret-slack-bot-token",
				Enabled:  true,
			},
			Discord: &struct {
				Webhook string "json:\"webhook\""
				Enabled bool   "json:\"enabled\""
			}{
				Webhook: "random",
				Enabled: true,
			},
			Webhook: &struct {
				Webhook string "json:\"webhook\""
				Enabled bool   "json:\"enabled\""
			}{
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
