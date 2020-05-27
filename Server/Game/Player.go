package Game

import (
	"github.com/gorilla/websocket"
	"log"
)

type Player struct {
	Name string
	Description string
	Connection *websocket.Conn
	CurrentRoom *Room
}

func (p Player) Inspect() string {
	return p.Description
}

func (p Player) SendMessage(message string) {
	err := p.Connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
	}
}

func (p Player) Logic() {
	for {
		mt, message, err := p.Connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			p.CurrentRoom.RemovePlayer(p.Name)
			break
		}
		log.Printf("recv: %s", message)
		err = p.Connection.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
		}
	}
}
