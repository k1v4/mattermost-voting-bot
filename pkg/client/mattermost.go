package client

import (
	"encoding/json"
	"fmt"
	"github.com/k1v4/mattermost-voting-bot/pkg/database/tarantoolDB"
	"github.com/mattermost/mattermost-server/v6/model"
	"strings"
)

type VotingBot struct {
	client        *model.Client4
	webSocket     *model.WebSocketClient
	user          *model.User
	team          *model.Team
	tarantoolConn *tarantoolDB.DB
}

type VotingBotConfig struct {
	MattermostURL string `json:"mattermost_url" env:"MATTERMOST_URL"`
	BotToken      string `json:"bot_token" env:"BOT_TOKEN"`
}

func NewVotingBot(cfg VotingBotConfig, storage *tarantoolDB.DB) (*VotingBot, error) {

	mmURL := cfg.MattermostURL
	wsURL := strings.Replace(mmURL, "http", "ws", 1) + "websocket"
	fmt.Println(wsURL)

	client := model.NewAPIv4Client(wsURL)
	client.SetToken(cfg.BotToken)

	webSocket, err := model.NewWebSocketClient4(wsURL, client.AuthToken)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения WebSocket: %w", err)
	}

	bot := &VotingBot{
		client:        client,
		webSocket:     webSocket,
		tarantoolConn: storage,
	}

	go bot.listenToMessages()

	return bot, nil
}

func (b *VotingBot) listenToMessages() {
	b.webSocket.Listen()
	for event := range b.webSocket.EventChannel {
		if event.EventType() == model.WebsocketEventPosted {
			// Получаем JSON строки из события
			postData, ok := event.GetData()["post"].(string)
			if !ok {
				continue
			}

			var post model.Post
			if err := json.NewDecoder(strings.NewReader(postData)).Decode(&post); err != nil {
				continue
			}

			if post.UserId == b.user.Id {
				continue
			}

			b.handleCommand(&post)
		}
	}
}

func (b *VotingBot) handleCommand(post *model.Post) {
	commandParts := strings.Fields(post.Message)
	if len(commandParts) == 0 {
		return
	}

	switch commandParts[0] {
	case "/create_vote":
		//b.createVote(post, commandParts[1:])
	case "/vote":
		//b.vote(post, commandParts[1:])
	case "/results":
		//b.getResults(post, commandParts[1:])
	default:
		//b.sendMessage(post.ChannelId, "Неизвестная команда. Доступные команды: /create_vote, /vote, /results")
	}
}
