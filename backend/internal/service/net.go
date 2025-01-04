package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService

	host *websocket.Conn
	tick int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (x *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case 0:
		{
			return ConnectPacket{}
		}
	}

	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case ConnectPacket:
		{
			return 0, nil
		}
	}

	return 0, errors.New("invalid packet type")
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

	fmt.Println(packetId)
	fmt.Println(packet)
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
