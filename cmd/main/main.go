package main

import (
	"context"
	"github.com/k1v4/mattermost-voting-bot/internal/config"
	"github.com/k1v4/mattermost-voting-bot/pkg/database/tarantoolDB"
	"github.com/k1v4/mattermost-voting-bot/pkg/logger"
	"github.com/tarantool/go-tarantool/v2"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
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

	votingBotLogger.Info(ctx, "read config successfully")

	storage, err := tarantoolDB.New(cfg.DBConfig)
	if err != nil {
		panic(err)
	}

}
