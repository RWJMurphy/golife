package conway

import "fmt"

type Game struct {
	Prev, Current *world
	width, height int
}

func NewGame(width, height int) *Game {
	return &Game{
		nil,
		NewWorld(width, height),
		width, height,
	}
}

func (g *Game) Next() {
	g.Prev, g.Current = g.Current, g.Current.Next()
}

func (g *Game) Print() {
	for _, row := range g.Current.cells {
		for _, cell := range row {
			if cell.alive {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
