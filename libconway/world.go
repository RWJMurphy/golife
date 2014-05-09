package conway

import "runtime"

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
	if x < 0 {
		x = w.width + x
	} else if x >= w.width {
		x = x % w.width
	}

	if y < 0 {
		y = w.height + y
	} else if y >= w.height {
		y = y % w.height
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

type update_job struct {
	w    *world
	x, y int
}

type update_result struct {
	x, y int
	cell Cell
}

func cell_update_worker(jobs <-chan update_job, results chan<- update_result) {
	for j := range jobs {
		cell := j.w.GetCell(j.x, j.y)
		neighbours := j.w.GetAliveNeighbours(j.x, j.y)
		next_cell := cell.Next(neighbours)
		results <- update_result{j.x, j.y, next_cell}
	}
}

func (w *world) Next() *world {
	jobs := make(chan update_job, 100)
	results := make(chan update_result, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		go cell_update_worker(jobs, results)
	}

	for y, row := range w.cells {
		for x := range row {
			jobs <- update_job{w, x, y}
		}
	}
	close(jobs)

	next := NewWorld(w.width, w.height)
	for i := 0; i < w.width*w.height; i++ {
		result := <-results
		next.cells[result.y][result.x] = result.cell
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
