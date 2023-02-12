package gameoflife

// NewCell constructs a Cell instance
func NewCell(isAlive bool) Cell {
	return Cell{alive: isAlive}
}

// Cell models a Cell in the GameOfLife Universe that can be in either
// Alive or Dead state and can evolve based on interactions with its
// neighbooring Cells
type Cell struct {
	alive bool
}

// IsAlive returns true if the Cell is considered to be in Alive state
func (c *Cell) IsAlive() bool {
	return c.alive == true
}

// Interact enforces the GameOfLife rules to a Cell and produces a new Cell with
// its state determined by the number of Alive neighbooring Cells
func (c *Cell) Interact(cells Cells) Cell {
	liveNeighboors := cells.TotalAlive()
	shouldLive := c.IsAlive()
	if c.IsAlive() {
		// Rule 1. Any live cell with fewer than two live neighbors dies
		// Rule 3. Any live cell with more than three live neighbors dies
		if liveNeighboors < 2 || liveNeighboors > 3 {
			shouldLive = false
		}
		// (all other cases) Rule 2. Any live cell with two or three live neighbors
		// lives on to the next generation.
	} else {
		// Rule 4. Any dead cell with exactly three live neighbors becomes a live cell
		if liveNeighboors == 3 {
			shouldLive = true
		}
	}
	return NewCell(shouldLive)
}

// Cells is a collection of Cell items
type Cells []Cell

// TotalAlive returns the total number of Alive Cells in the collection
func (n *Cells) TotalAlive() int {
	count := 0
	for _, cell := range *n {
		if cell.IsAlive() {
			count += 1
		}
	}
	return count
}
