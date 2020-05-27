package Game

import "errors"

type GameManager struct {
	Players map[string] *Player
	Updateables []Updateable
	StartingRoom Room
}

func (g GameManager) Update()  {
	for _, u := range g.Updateables {
		u.Update()
	}
}

func (g GameManager) AddPlayer(name string, player *Player) error {
	if g.Players[name] != nil {
		return errors.New("Player in use")
	}
	g.Players[name] = player
	return nil
}

func (g GameManager) RemovePlayer(name string) error {
	if g.Players[name] == nil {
		return errors.New("Player not found")
	}
	g.Players[name].CurrentRoom.RemovePlayer(g.Players[name].Name)
	delete(g.Players, name)
	return nil
}
