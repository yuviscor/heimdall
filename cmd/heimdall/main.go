package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/MowlCoder/heimdall/internal/checker"
	"github.com/MowlCoder/heimdall/internal/config"
	"github.com/MowlCoder/heimdall/internal/notifier"
)

func main() {
	flagCfg := config.NewFlagConfig()
	if err := flagCfg.Parse(); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfigFromFile(flagCfg.PathToConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	telegramNotifier := notifier.NewTelegramNotifier(
		cfg.Telegram.ChatID,
		cfg.Telegram.BotToken,
	)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	serviceChecker := checker.NewServiceChecker(telegramNotifier, cfg.Services)
	serviceChecker.Start(ctx)

	log.Println("[INFO]: heimdall service status checker was started")

	<-ctx.Done()
	stop()

	log.Println("[INFO]: starting shutdown of status checker...")

	serviceChecker.WaitShutdown()
}
