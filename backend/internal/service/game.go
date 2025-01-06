package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"backend/internal/entity"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Player struct {
	Id         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Connection *websocket.Conn `json:"-"`
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

type Game struct {
	Id      uuid.UUID
	Quiz    entity.Quiz
	Code    string
	State   GameState
	Players []Player

	Host       *websocket.Conn
	NetService *NetService
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func newGame(quiz entity.Quiz, host *websocket.Conn, netService *NetService) Game {
	return Game{
		Id:         uuid.New(),
		Quiz:       quiz,
		Code:       generateCode(),
		State:      LobbyState,
		Players:    []Player{},
		Host:       host,
		NetService: netService,
	}
}

func (g *Game) Start() {
	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {
	fmt.Println("tick")
}

func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	fmt.Println(name, "joined the game")

	player := Player{
		Id:         uuid.New(),
		Name:       name,
		Connection: connection,
	}

	g.Players = append(g.Players, player)

	g.NetService.SendPacket(connection, ChangeGameStatePacket{
		State: g.State,
	})

	g.NetService.SendPacket(g.Host, PlayerJoinPacket{
		Player: player,
	})
}
