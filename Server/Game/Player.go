package Game

import (
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

type Player struct {
	Name        string
	Description string
	Connection  *websocket.Conn
	CurrentRoom *Room
	GMan        *GameManager
}

func (p *Player) Inspect() string {
	return p.Description
}

func (p *Player) SendMessage(message string) {
	err := p.Connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
	}
}

func (p *Player) Logic() {
	for {
		_, message, err := p.Connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			p.CurrentRoom.RemovePlayer(p.Name)
			break
		}
		f := strings.Split(string(message), "::")
		var mDir Direction
		if len(f) > 1 {
			switch strings.ToLower(f[0]) {
			case "m":
			case "move":
				p.SendMessage("Which way are you going? [n]orth [e]ast [s]outh [w]est [u]p [d]own")
				continue
			case "n":
				fallthrough
			case "north":
				mDir = North
				break
			case "e":
				fallthrough
			case "east":
				mDir = East
				break
			case "s":
				fallthrough
			case "south":
				mDir = South
				break
			case "w":
				fallthrough
			case "west":
				mDir = West
				break
			case "u":
				fallthrough
			case "up":
				mDir = Up
				break
			case "d":
				fallthrough
			case "down":
				mDir = Down
				break
			}
		}
		log.Println("moving " + mDir.String())
		if val, ok := p.CurrentRoom.Exits[mDir]; ok {
			log.Println(val)
			if rVal, rOk := p.GMan.Rooms[val]; rOk {
				log.Println("entering " + rVal.Name)
				rVal.Enter(p.CurrentRoom, p)
			}
		}
	}
}
