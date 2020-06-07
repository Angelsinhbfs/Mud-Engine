package Game

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Up
	Down
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West", "Up", "Down"}[d]
}
