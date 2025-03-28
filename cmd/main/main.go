package main

import (
	"context"
	"github.com/k1v4/mattermost-voting-bot/internal/config"
	"github.com/k1v4/mattermost-voting-bot/pkg/client"
	"github.com/k1v4/mattermost-voting-bot/pkg/database/tarantoolDB"
	"github.com/k1v4/mattermost-voting-bot/pkg/logger"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
	"os"
	"os/signal"
	"syscall"
)

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	dialer := tarantool.NetDialer{
//		Address:  "127.0.0.1:3301",
//		User:     "vote_bot",
//		Password: "password",
//	}
//	opts := tarantool.Opts{
//		Timeout: time.Second,
//	}
//
//	conn, err := tarantool.Connect(ctx, dialer, opts)
//	if err != nil {
//		fmt.Println("Connection refused:", err)
//		return
//	}
//
//	fmt.Println("Connection established ", conn.Addr())
//}

func main() {
	ctx := context.Background()
	votingBotLogger := logger.NewLogger()
	ctx = context.WithValue(ctx, logger.LoggerKey, votingBotLogger)

	cfg := config.MustLoadConfig()
	if cfg == nil {
		panic("load config fail")
	}
	votingBotLogger.Info(ctx, "Config loaded successfully")

	storage, err := tarantoolDB.New(cfg.DBConfig)
	if err != nil {
		panic(err)
	}
	votingBotLogger.Info(ctx, "Connected to Tarantool")

	_, err = client.NewVotingBot(client.VotingBotConfig{
		MattermostURL: cfg.MattermostURL,
		BotToken:      cfg.BotToken,
	}, storage)
	if err != nil {
		panic(err)
	}
	votingBotLogger.Info(ctx, "Voting bot started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	votingBotLogger.Info(ctx, "Shutting down bot...")
}
