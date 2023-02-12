package gameoflife

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// NewUniverse constructs a Universe of the given size containing
// Cells on their default state (dead)
func NewUniverse(width int, height int) Universe {
	u := make([][]Cell, height)
	for i := range u {
		u[i] = make([]Cell, width)
	}
	return u
}

// NewRandomUniverse produces a randomized Universe of the given size
func NewRandomUniverse(width int, height int) Universe {
	rand.Seed(time.Now().Unix())
	u := NewUniverse(width, height)
	for y := 0; y < len(u); y++ {
		for x := 0; x <= len(u[y]); x++ {
			u.Set(x, y, rand.Intn(2) == 1)
		}
	}
	return u
}

// NewGliderUniverse produces a 25x25 Universe containing the Glider pattern
func NewGliderUniverse() Universe {
	size := 25
	u := NewUniverse(size, size)
	center := size / 2
	u.Set(center, center, true)
	u.Set(center+1, center, true)
	u.Set(center+2, center, true)
	u.Set(center+1, center-2, true)
	u.Set(center+2, center-1, true)
	return u
}

// Universe models the GameOfLife Universe as a 2D matrix of Cells.
// Although the Universe is in theory infinite, in practise we are only looking
// at a specific part of it at a time, hence it has a certain size
type Universe [][]Cell

// Clone creates a deep copy of the Universe
func (u *Universe) Clone() Universe {
	cells := *u
	cloned := make(Universe, len(cells))
	for y := 0; y < len(cells); y++ {
		cloned[y] = make([]Cell, len(cells[y]))
		for x := 0; x < len(cells[y]); x++ {
			cloned[y][x] = cells[y][x]
		}
	}
	return cloned
}

// Get a Cell at the given position. Since the universe is infinite, asking for
// a Cell outside the matrix boundaries will (safely) return a dead Cell
func (u *Universe) Get(x int, y int) *Cell {
	cells := *u
	if y >= 0 && y < len(cells) {
		if x >= 0 && x < len(cells[y]) {
			return &cells[y][x]
		}
	}
	return &Cell{false}
}

// Set the state of a Cell at the given position
func (u *Universe) Set(x int, y int, alive bool) error {
	cells := *u
	if y < len(cells) && x < len(cells[y]) {
		cells[y][x] = Cell{alive}
		return nil
	}
	return fmt.Errorf("Position X:%d, Y:%d is outside the boundaries", x, y)
}

// GetNeighboors returns a list of Cells surrounding a Cell at a given position
func (u *Universe) GetNeighboors(x int, y int) Cells {
	neighboors := make(Cells, 8)
	index := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// exclude current cell
			if i == 0 && j == 0 {
				continue
			}
			if x+j >= 0 && y+i >= 0 {
				cell := u.Get(x+j, y+i)
				neighboors[index] = *cell
				index++
			}
		}
	}
	return neighboors
}

// Evolve produces the next generation of a universe by triggering Cell
// interactions
func (u *Universe) Evolve() Universe {
	cloned := u.Clone()
	for y := 0; y < len(cloned); y++ {
		for x := 0; x <= len(cloned[y]); x++ {
			cell := u.triggerInteraction(x, y)
			cloned.Set(x, y, cell.IsAlive())
		}
	}
	return cloned
}

// triggerInteraction will have the Cell at the given position interact with
// its neighboors and produce an evolved Cell
func (u *Universe) triggerInteraction(x int, y int) Cell {
	cell := u.Get(x, y)
	neighboors := u.GetNeighboors(x, y)
	return cell.Interact(neighboors)
}

// ToString serializes the Universe into a string
func (u *Universe) ToString() string {
	builder := strings.Builder{}
	cells := *u
	for y := 0; y < len(cells); y++ {
		for x := 0; x < len(cells[y]); x++ {
			if u.Get(x, y).IsAlive() {
				builder.WriteString("*")
				continue
			}
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
