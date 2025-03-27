package client

import (
	"github.com/k1v4/mattermost-voting-bot/internal/config"
	"github.com/k1v4/mattermost-voting-bot/pkg/database/tarantoolDB"
	"github.com/mattermost/mattermost-server/v6/model"
)

type VotingBot struct {
	client        *model.Client4
	webSocket     *model.WebSocketClient
	user          *model.User
	team          *model.Team
	tarantoolConn tarantoolDB.DB
}

type VotingBotConfig struct {
	MattermostURL string `json:"mattermost_url,omitempty"`
	BotToken      string `json:"bot_token,omitempty"`
}

func NewVotingBot(cfg VotingBotConfig, storage tarantoolDB.DB) (*VotingBot, error) {
	client := model.NewAPIv4Client(cfg.MattermostURL)
	client.SetToken(cfg.BotToken)

	// Создание экземпляра бота
	bot := &VotingBot{
		client:        client,
		tarantoolConn: storage,
	}

	return bot, nil
}
