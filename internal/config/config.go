package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type Config struct {
	Services  []domain.Service `json:"services"`
	Notifiers struct {
		Telegram *struct {
			ChatID   string `json:"chatId"`
			BotToken string `json:"botToken"`
			Enabled  bool   `json:"enabled"`
		}
		Discord *struct {
			Webhook string `json:"webhook"`
			Enabled bool   `json:"enabled"`
		}
		Slack *struct {
			ChatID   string `json:"chatId"`
			BotToken string `json:"botToken"`
			Enabled  bool   `json:"enabled"`
		}
	} `json:"notifiers"`
}

func (c Config) IsTelegramEnabled() bool {
	return c.Notifiers.Telegram != nil && c.Notifiers.Telegram.Enabled
}

func (c Config) IsDiscordEnabled() bool {
	return c.Notifiers.Discord != nil && c.Notifiers.Discord.Enabled
}

func (c Config) IsSlackEnabled() bool {
	return c.Notifiers.Slack != nil && c.Notifiers.Slack.Enabled
}

func ParseConfigFromFile(path string) (*Config, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := json.Unmarshal(fileContent, &cfg); err != nil {
		return nil, err
	}

	if !cfg.IsTelegramEnabled() && !cfg.IsDiscordEnabled() && !cfg.IsSlackEnabled() {
		return nil, errors.New("at least 1 notifier service should be enabled and configured")
	}

	return &cfg, nil
}
