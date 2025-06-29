<p align="center">
  <img src="./logo.png" alt="Heimdall Logo" width="150"">
</p>

<h2 align="center">
  🔍 Heimdall - Service Monitoring Guardian 🔍
</h2>

<p align="center">
  <em>Watchful protector of your services</em>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.24.2+-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://goreportcard.com/badge/github.com/MowlCoder/heimdall" alt="Go Report Card">
  <img src="https://img.shields.io/badge/Dependencies-None-green.svg" alt="Dependencies">

  <hr>
</p>

**Heimdall** - A lightweight, high-performance service monitoring tool that continuously checks your services' health and sends instant notifications when they become unavailable.

## ✨ Features

- **🔄 Continuous Monitoring**: Real-time health checks with configurable intervals
- **📱 Multi-Platform Notifications**: Instant alerts via Telegram, Discord or Slack
- **⚡ High Performance**: Lightweight Go implementation with minimal resource usage
- **🎯 Flexible Configuration**: Support for both strict and lenient health checks
- **🚀 Zero Dependencies**: No external dependencies - only Go standard library

## 🚀 Quick Start

### Prerequisites

- Go 1.24.2 or higher
- Telegram Bot Token (for Telegram notifications)
- Discord Webhook URL (for Discord notifications)
- Slack Bot Token (for Slack notifications)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/MowlCoder/heimdall.git
   cd heimdall
   ```

2. **Build the application**
   ```bash
   # Using Task (recommended)
   task build
   
   # Or using Go directly
   go build -o bin/heimdall ./cmd/heimdall/main.go
   ```

3. **Create configuration file**
   ```bash
   cp config.example.json config.json
   ```

4. **Configure your services and notifiers**
   ```json
   {
     "services": [
       {
         "name": "My Website",
         "url": "https://example.com",
         "interval": "30s",
         "strict": true,
         "headers": {
           "User-Agent": "Heimdall-Monitor/1.0",
           "Authorization": "Bearer your-token"
         }
       },
       {
         "name": "API Service",
         "url": "https://api.example.com/health",
         "interval": "1m",
         "strict": false,
         "timeout": 5000
       }
     ],
     "notifiers": {
       "telegram": {
         "chatId": "YOUR_CHAT_ID",
         "botToken": "YOUR_BOT_TOKEN",
         "enabled": true
       },
       "discord": {
         "webhook": "YOUR_DISCORD_WEBHOOK_URL",
         "enabled": false
       },
       "slack": {
         "chatId": "YOUR_CHANNEL_ID",
         "botToken": "YOUR_SLACK_BOT_TOKEN",
         "enabled": true
       }
     }
   }
   ```

5. **Run the service**
   ```bash
   # Using Task
   task run
   
   # Or directly
   ./bin/heimdall --config ./config.json
   ```

## 📋 Configuration

### Service Configuration

Each service in the `services` array supports the following parameters:

| Parameter | Type | Required | Description | Example |
|-----------|------|----------|-------------|---------|
| `name` | string | ✅ | Human-readable service name | `"My API"` |
| `url` | string | ✅ | Service URL to monitor | `"https://api.example.com"` |
| `interval` | string | ✅ | Check interval (supports Go duration or milliseconds) | `"30s"`, `"1m"`, `"30000"` |
| `strict` | boolean | ❌ | Strict mode (only 200 OK considered healthy) | `true` |
| `timeout` | integer | ❌ | Request timeout in milliseconds | `600` |
| `headers` | object | ❌ | Custom HTTP headers to send with requests | `{"Authorization": "Bearer token"}` |

### Notifiers Configuration

#### Telegram Notifier

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `chatId` | string | ✅ | Telegram chat ID where notifications will be sent |
| `botToken` | string | ✅ | Telegram bot token from @BotFather |
| `enabled` | boolean | ❌ | Enable/disable Telegram notifications (default: `true`) |

#### Discord Notifier

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `webhook` | string | ✅ | Discord webhook URL for the channel |
| `enabled` | boolean | ❌ | Enable/disable Discord notifications (default: `true`) |

#### Slack Notifier

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `chatId` | string | ✅ | Slack channel ID where notifications will be sent |
| `botToken` | string | ✅ | Slack bot token |
| `enabled` | boolean | ❌ | Enable/disable Slack notifications (default: `true`) |

### Health Check Modes

#### Strict Mode (`strict: true`)
- Only HTTP 200 status code is considered healthy
- Any other status code triggers a notification
- Ideal for APIs that should always return 200 OK

#### Lenient Mode (`strict: false` or omitted)
- HTTP 200-499 status codes are considered healthy
- Only 5xx server errors trigger notifications
- Suitable for services that may return redirects or client errors

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Named after Heimdall, the Norse god who watches over the Bifröst bridge
- Built with Go for performance and reliability
- Telegram Bot API, Discord Webhooks, and Slack API for instant notifications
- Rich message formatting with emojis, timestamps, and structured information

---

**Made with ❤️ for reliable service monitoring**