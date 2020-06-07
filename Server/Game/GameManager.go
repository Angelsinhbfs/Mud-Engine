package Game

import "errors"

type GameManager struct {
	Players      map[string]*Player
	Updateables  []Updateable
	Rooms        map[string]*Room
	StartingRoom *Room
}

func (g *GameManager) Update() {
	for _, u := range g.Updateables {
		go u.Update()
	}
}

func (g *GameManager) AddPlayer(name string, player *Player) error {
	if g.Players[name] != nil {
		return errors.New("Player in use")
	}
	g.Players[name] = player
	return nil
}

func (g *GameManager) RemovePlayer(name string) error {
	if g.Players[name] == nil {
		return errors.New("Player not found")
	}
	g.Players[name].CurrentRoom.RemovePlayer(g.Players[name].Name)
	delete(g.Players, name)
	return nil
}

func (g *GameManager) AddRoom(room *Room) {
	if room.UID == "Default" {
		g.StartingRoom = room
	}
	g.Updateables = append(g.Updateables, room)
	g.Rooms[room.UID] = room
	g.Players = make(map[string]*Player)
}
