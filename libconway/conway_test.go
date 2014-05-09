package conway

import (
	"fmt"
	"testing"
)

func Expect(condition interface{}, expected interface{}, message string, t *testing.T) {
	if condition != expected {
		t.Error(fmt.Sprintf("%s; expected %+v, got %+v", message, expected, condition))
	}
}

func TestCellAliveRule(t *testing.T) {
	Expect(NewCell(true).Next(1).alive, false, "Alive cell with 1 neighbour should die", t)
	Expect(NewCell(true).Next(2).alive, true, "Alive cell with 2 neighbour should live", t)
	Expect(NewCell(true).Next(3).alive, true, "Alive cell with 3 neighbour should live", t)
	Expect(NewCell(true).Next(4).alive, false, "Alive cell with 4 neighbour should die", t)
}

func TestCellDeadRule(t *testing.T) {
	Expect(NewCell(false).Next(1).alive, false, "Dead cell with 1 neighbour should stay dead", t)
	Expect(NewCell(false).Next(2).alive, false, "Dead cell with 2 neighbour should stay dead", t)
	Expect(NewCell(false).Next(3).alive, true, "Dead cell with 3 neighbour should come alive", t)
	Expect(NewCell(false).Next(4).alive, false, "Dead cell with 4 neighbour should stay dead", t)
}

func TestCellIsImmutable(t *testing.T) {
	dead_cell := NewCell(false)
	dead_cell_addr := &dead_cell
	alive_cell := NewCell(true)
	alive_cell_addr := &alive_cell
	for x := 0; x < 9; x++ {
		dead_cell.Next(x)
		alive_cell.Next(x)
		Expect(alive_cell.alive, true, "Alive cell changed state", t)
		Expect(&alive_cell, alive_cell_addr, "Alive cell address changed", t)
		Expect(dead_cell.alive, false, "Dead cell changed state", t)
		Expect(&dead_cell, dead_cell_addr, "Dead cell address changed", t)
	}
}

func TestGameSetup(t *testing.T) {
	world := NewWorld(5, 5)
	world.SetCellAlive(1, 3)
	Expect(world.GetCell(1, 3).alive, true, "Cell died without any change", t)
	Expect(world.GetCell(1, 4).alive, false, "Cell cannot be alive now", t)
}

func TestNeighbourAliveCount(t *testing.T) {
	world := NewWorld(5, 5)

	world.SetCellAlive(1, 2)
	world.SetCellAlive(2, 2)
	world.SetCellAlive(3, 2)

	Expect(world.GetAliveNeighbours(0, 0), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(1, 0), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(2, 0), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(3, 0), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(4, 0), 0, "Incorrect alive neighbour count", t)

	Expect(world.GetAliveNeighbours(0, 1), 1, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(1, 1), 2, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(2, 1), 3, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(3, 1), 2, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(4, 1), 1, "Incorrect alive neighbour count", t)

	Expect(world.GetAliveNeighbours(0, 2), 1, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(1, 2), 1, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(2, 2), 2, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(3, 2), 1, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(4, 2), 1, "Incorrect alive neighbour count", t)

	Expect(world.GetAliveNeighbours(0, 3), 1, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(1, 3), 2, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(2, 3), 3, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(3, 3), 2, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(4, 3), 1, "Incorrect alive neighbour count", t)

	Expect(world.GetAliveNeighbours(0, 4), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(1, 4), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(2, 4), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(3, 4), 0, "Incorrect alive neighbour count", t)
	Expect(world.GetAliveNeighbours(4, 4), 0, "Incorrect alive neighbour count", t)

}

func TestWorldNext(t *testing.T) {
	world := NewWorld(5, 5)

	world.SetCellAlive(1, 2)
	world.SetCellAlive(2, 2)
	world.SetCellAlive(3, 2)

	next_world := world.Next()

	Expect(next_world.GetCell(2, 1).alive, true, "Incorrect cell state at 2, 1", t)
	Expect(next_world.GetCell(2, 2).alive, true, "Incorrect cell state at 2, 2", t)
	Expect(next_world.GetCell(2, 3).alive, true, "Incorrect cell state at 2, 3", t)

	Expect(next_world.GetCell(1, 2).alive, false, "Incorrect cell state at 1, 2", t)
	Expect(next_world.GetCell(3, 2).alive, false, "Incorrect cell state at 3, 2", t)
}
