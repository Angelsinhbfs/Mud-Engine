package Game

type Room struct {
	Name string
	Description string
	Exits map[Direction] *Room
	Players map[string] *Player
}

func (r Room) Enter (oldRoom *Room, player *Player) {
	if oldRoom != nil {
		delete(oldRoom.Players, player.Name)
	}
	r.SendMessage(player.Name + " has entered the room")
	r.Players[player.Name] = player
	player.SendMessage("You have entered " + r.Name + ". When you look around you see " + r.Description)
	player.CurrentRoom = &r
}

func (r Room) SendMessage(message string) {
	for _, p := range r.Players {
		p.SendMessage(message)
	}
}

func (r Room) Update() {

}

func (r Room) RemovePlayer(name string) {
	if r.Players[name] == nil {
		return
	}
	r.SendMessage(name  + " disappears in a puff of smoke")
	delete(r.Players, name)
}

