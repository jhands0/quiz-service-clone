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
	Answered   bool            `json:"-"`
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

type Game struct {
	Id              uuid.UUID
	Quiz            entity.Quiz
	CurrentQuestion int
	Code            string
	State           GameState
	Time            int
	Players         []*Player

	Host       *websocket.Conn
	NetService *NetService
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func newGame(quiz entity.Quiz, host *websocket.Conn, netService *NetService) Game {
	return Game{
		Id:              uuid.New(),
		Quiz:            quiz,
		CurrentQuestion: -1,
		Code:            generateCode(),
		State:           LobbyState,
		Time:            60,
		Players:         []*Player{},

		Host:       host,
		NetService: netService,
	}
}

func (g *Game) Start() {
	g.ChangeState(PlayState)
	g.NextQuestion()

	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) NextQuestion() {
	g.CurrentQuestion++

	g.ChangeState(PlayState)
	g.Time = 60

	g.NetService.SendPacket(g.Host, QuestionShowPacket{
		Question: g.Quiz.Questions[g.CurrentQuestion],
	})
}

func (g *Game) Reveal() {
	g.Time = 10
	g.ChangeState(RevealState)
}

func (g *Game) Tick() {
	g.Time--
	g.NetService.SendPacket(g.Host, TickPacket{
		Tick: g.Time,
	})

	if g.Time == 0 {
		switch g.State {
		case PlayState:
			{
				g.Reveal()
				break
			}
		case RevealState:
			{
				g.NextQuestion()
				break
			}
		}
	}
}

func (g *Game) ChangeState(state GameState) {
	g.State = state
	g.BroadcastPacket(ChangeGameStatePacket{
		State: state,
	}, true)
}

func (g *Game) BroadcastPacket(packet any, includeHost bool) error {
	for _, player := range g.Players {
		err := g.NetService.SendPacket(player.Connection, packet)
		if err != nil {
			return err
		}
	}

	if includeHost {
		err := g.NetService.SendPacket(g.Host, packet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	fmt.Println(name, "joined the game")

	player := Player{
		Id:         uuid.New(),
		Name:       name,
		Connection: connection,
	}

	g.Players = append(g.Players, &player)

	g.NetService.SendPacket(connection, ChangeGameStatePacket{
		State: g.State,
	})

	g.NetService.SendPacket(g.Host, PlayerJoinPacket{
		Player: player,
	})
}

func (g *Game) getAnsweredPlayers() []*Player {
	players := []*Player{}

	for _, player := range g.Players {
		if player.Answered {
			players = append(players, player)
		}
	}

	return players
}

func (g *Game) OnPlayerAnswer(question int, player *Player) {
	player.Answered = true

	if len(g.getAnsweredPlayers()) == len(g.Players) {
		fmt.Println("debug: all players have answered!")
		g.Reveal()
	}
}
