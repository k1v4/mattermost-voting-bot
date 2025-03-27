package tarantoolDB

import (
	"context"
	"fmt"
	"github.com/tarantool/go-tarantool/v2"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
	"log"
	"time"
)

type DBConfig struct {
	Host     string        `json:"host"`
	Port     int           `json:"port"`
	User     string        `json:"user"`
	Password string        `json:"password"`
	Timeout  time.Duration `json:"timeout"`
}

type DB struct {
	Db *tarantool.Connection
}

func New(cfg DBConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dialer := tarantool.NetDialer{
		Address:  fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
	}

	opts := tarantool.Opts{
		Timeout: cfg.Timeout,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		log.Fatalf("Connection refused: %v", err)
	}

	return &DB{
		Db: conn,
	}, err
}
