package conway

type Coord struct {
	x, y int
}

type world struct {
	width, height int
	cells         [][]Cell
}

func NewWorld(width, height int) *world {
	cells := make([][]Cell, height)
	for y := range cells {
		cells[y] = make([]Cell, width)
	}
	return &world{width, height, cells}
}

func (w *world) SetCellAlive(x, y int) {
	w.cells[y][x] = NewCell(true)
}

func (w *world) GetCell(x, y int) Cell {
	if x < 0 || y < 0 || x >= w.width || y >= w.height {
		return NewCell(false)
	}
	return w.cells[y][x]
}

var offsets = [][]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (w *world) GetAliveNeighbours(x, y int) int {
	n_count := 0
	for _, offset := range offsets {
		if w.GetCell(x+offset[0], y+offset[1]).alive {
			n_count++
		}
	}
	return n_count
}

func (w *world) Next() *world {
	next := NewWorld(w.width, w.height)
	for y, row := range w.cells {
		for x, cell := range row {
			next.cells[y][x] = cell.Next(w.GetAliveNeighbours(x, y))
		}
	}
	return next
}

func (w world) Equal(other *world) bool {
	for y, row := range w.cells {
		for x, cell := range row {
			if cell.alive != other.GetCell(x, y).alive {
				return false
			}
		}
	}
	return true
}
