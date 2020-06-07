package Game

import "log"

type Room struct {
	UID         string
	Name        string
	Description string
	Exits       map[Direction]string //UID of connecting room
	Players     map[string]*Player
	Items       map[string]*Item
}

func (r *Room) New() {
	r.Exits = make(map[Direction]string)
	r.Players = make(map[string]*Player)
	r.Items = make(map[string]*Item)
}

func (r *Room) Enter(oldRoom *Room, player *Player) {
	if oldRoom != nil {
		delete(oldRoom.Players, player.Name)
	}
	r.SendMessage("sys::" + player.Name + " has entered the room")
	r.Players[player.Name] = player
	player.SendMessage("d::" + "You have entered " + r.Name + ". When you look around you see " + r.Description)
	player.CurrentRoom = r
}

func (r *Room) SendMessage(message string) {
	for _, p := range r.Players {
		p.SendMessage(message)
	}
}

func (r *Room) Update() {
	log.Println("Tick " + r.Name + " updated")
}

func (r *Room) RemovePlayer(name string) {
	if r.Players[name] == nil {
		return
	}
	r.SendMessage("sys::" + name + " disappears in a puff of smoke")
	delete(r.Players, name)
}
