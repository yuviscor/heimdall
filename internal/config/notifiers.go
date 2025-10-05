package config

type TelegramNotifierConfig struct {
	ChatID   string `json:"chatId"`
	BotToken string `json:"botToken"`
	Enabled  bool   `json:"enabled"`
}

type DiscordNotifierConfig struct {
	Webhook string `json:"webhook"`
	Enabled bool   `json:"enabled"`
}

type SlackNotifierConfig struct {
	ChatID   string `json:"chatId"`
	BotToken string `json:"botToken"`
	Enabled  bool   `json:"enabled"`
}

type WebhookNotifierConfig struct {
	Webhook string `json:"webhook"`
	Enabled bool   `json:"enabled"`
}
