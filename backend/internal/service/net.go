package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"backend/internal/entity"
	//"backend/internal/game"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/internal/uuid"
)

type NetService struct {
	quizService *QuizService

	games []*Game
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
		games:       []*Game{},
	}
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HostGamePacket struct {
	QuizId string `json:"quizId"`
}

type QuestionShowPacket struct {
	Question entity.QuizQuestion `json:"question"`
}

type ChangeGameStatePacket struct {
	State GameState `json:"state"`
}

type PlayerJoinPacket struct {
	Player Player `json:"player"`
}

type PlayerDisconnectPacket struct {
	PlayerId uuid.UUID `json:"playerId"`
}

type StartGamePacket struct {
}

type TickPacket struct {
	Tick int `json:"tick"`
}

type QuestionAnswerPacket struct {
	Question int `json:"question"`
}

type PlayerRevealPacket struct {
	Points int `json:"points"`
}

type LeaderboardPacket struct {
	Points []LeaderboardEntry `json:"points"`
}

func (c *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case 0:
		{
			return &ConnectPacket{}
		}
	case 1:
		{
			return &HostGamePacket{}
		}
	case 5:
		{
			return &StartGamePacket{}
		}
	case 7:
		{
			return &QuestionAnswerPacket{}
		}
	}

	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case HostGamePacket:
		{
			return 1, nil
		}
	case QuestionShowPacket:
		{
			return 2, nil
		}
	case ChangeGameStatePacket:
		{
			return 3, nil
		}
	case PlayerJoinPacket:
		{
			return 4, nil
		}
	case TickPacket:
		{
			return 6, nil
		}
	case PlayerRevealPacket:
		{
			return 8, nil
		}
	case LeaderboardPacket:
		{
			return 9, nil
		}
	case PlayerDisconnectPacket:
		{
			return 10, nil
		}
	}

	return 0, errors.New("invalid packet type")
}

func (c *NetService) getGameByCode(code string) *Game {
	for _, game := range c.games {
		if game.Code == code {
			return game
		}
	}

	return nil
}

func (c *NetService) getGameByHost(host *websocket.Conn) *Game {
	for _, game := range c.games {
		if game.Host == host {
			return game
		}
	}

	return nil
}

func (c *NetService) getGameByPlayer(con *websocket.Conn) (*Game, *Player) {
	for _, game := range c.games {
		for _, player := range game.Players {
			if player.Connection == con {
				return game, player
			}
		}
	}

	return nil, nil
}

func (c *NetService) OnDisconnect(con *websocket.Conn) {
	game, player := c.getGameByPlayer(con)
	if game == nil {
		return
	}

	game.OnPlayerDisconnect(player)
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 1 {
		return
	}

	packetId := msg[0]
	data := msg[1:]

	packet := c.packetIdToPacket(packetId)
	if packet == nil {
		return
	}

	err := json.Unmarshal(data, &packet)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch data := packet.(type) {
	case *ConnectPacket:
		{
			game := c.getGameByCode(data.Code)
			if game == nil {
				return
			}

			game.OnPlayerJoin(data.Name, con)
			break
		}
	case *HostGamePacket:
		{
			quizId, err := primitive.ObjectIDFromHex(data.QuizId)
			if err != nil {
				fmt.Println(err)
				return
			}

			quiz, err := c.quizService.quizCollection.GetQuizById(quizId)
			if err != nil {
				fmt.Println(err)
				return
			}

			if quiz == nil {
				return
			}

			newGame := newGame(*quiz, con, c)
			fmt.Println("new game: ", newGame.Code)
			c.games = append(c.games, &newGame)

			c.SendPacket(con, HostGamePacket{
				QuizId: newGame.Code,
			})

			c.SendPacket(con, ChangeGameStatePacket{
				State: newGame.State,
			})
			break
		}
	case *StartGamePacket:
		{
			game := c.getGameByHost(con)
			if game == nil {
				return
			}

			game.StartOrSkip()
			break
		}
	case *QuestionAnswerPacket:
		{
			game, player := c.getGameByPlayer(con)
			if game == nil {
				return
			}

			game.OnPlayerAnswer(data.Question, player)
			break
		}
	}
}

func (c *NetService) SendPacket(connection *websocket.Conn, packet any) error {
	bytes, err := c.PacketToBytes(packet)
	if err != nil {
		return err
	}

	return connection.WriteMessage(websocket.BinaryMessage, bytes)
}

func (c *NetService) PacketToBytes(packet any) ([]byte, error) {
	packetId, err := c.packetToPacketId(packet)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	final := append([]byte{packetId}, bytes...)
	return final, nil
}
