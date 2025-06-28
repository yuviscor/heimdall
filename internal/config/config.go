package config

import (
	"encoding/json"
	"os"

	"github.com/MowlCoder/heimdall/internal/domain"
)

type Config struct {
	Services []domain.Service `json:"services"`
	Telegram struct {
		ChatID   string `json:"chatId"`
		BotToken string `json:"botToken"`
	}
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

	return &cfg, nil
}
