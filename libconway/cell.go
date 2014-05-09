package conway

type Cell struct {
	alive bool
}

func NewCell(alive bool) Cell {
	return Cell{alive}
}

func (cell Cell) Next(neighbours int) Cell {
	return Cell{neighbours == 3 || (cell.alive && neighbours == 2)}
}
