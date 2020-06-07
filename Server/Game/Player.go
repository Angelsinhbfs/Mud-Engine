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
	Equipment   Equipment
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
		if len(f) > 0 {
			switch strings.ToLower(f[0]) {
			case "m", "move":
				p.SendMessage("Which way are you going? [n]orth [e]ast [s]outh [w]est [u]p [d]own")
				continue
			case "n", "north", "e", "east", "s", "south", "w", "west", "u", "up", "d", "down":
				p.Move(strings.ToLower(f[0]))
				continue
			case "l":
				fallthrough
			case "look":
				p.SendMessage("d::" + p.CurrentRoom.Description)
				continue
			case "p", "pick up", "i", "inventory", "eq", "equip", "uq", "unequip", "dr", "drop":
				p.inventoryAction(f)
			case "a":
				fallthrough
			case "attack":
				p.SendMessage("sys::Not yet implemented")
				continue

			case "wh":
				fallthrough
			case "whisper":
				p.SendMessage("sys::Not yet implemented")
				continue
			default:
				p.CurrentRoom.SendMessage("s::" + p.Name + " says: " + string(message))
				continue
			}
		} else {
			p.CurrentRoom.SendMessage("s::" + p.Name + " says: " + string(message))
		}

	}
}

func (p *Player) Move(d string) {
	var mDir Direction
	switch d {
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
	if val, ok := p.CurrentRoom.Exits[mDir]; ok {
		log.Println(val)
		if rVal, rOk := p.GMan.Rooms[val]; rOk {
			rVal.Enter(p.CurrentRoom, p)
		}
	} else {
		p.SendMessage("You cannot got that way")
	}
}

func (p *Player) inventoryAction(i []string) {
	switch i[0] {
	case "i", "inventory":
		p.SendMessage("sys::Not yet implemented")
		break
	case "eq", "equip":
		p.SendMessage("sys::Not yet implemented")
		break
	case "uq", "unequip":
		p.SendMessage("sys::Not yet implemented")
		break
	case "p", "pick up":
		p.SendMessage("sys::Not yet implemented")
		break
	case "dr", "drop":
		p.SendMessage("sys::Not yet implemented")
		break
	}
}
