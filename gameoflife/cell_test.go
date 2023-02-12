package gameoflife

import "testing"

type TestCace struct {
	startAlive bool
	neighboors Cells
	endAlive   bool
}

func TestRule1CellUnderpopulation(t *testing.T) {
	cases := []TestCace{
		{true, Cells{{true}, {false}}, false},
		{true, Cells{{true}, {true}, {false}}, true},
	}

	for _, c := range cases {
		t.Run("Testing Rule 1", func(t *testing.T) {
			cell := NewCell(c.startAlive)
			newcell := cell.Interact(c.neighboors)
			if newcell.IsAlive() != c.endAlive {
				t.Errorf("Rule 1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation. Cell Alive?: %v, Neighboors: %d", newcell.IsAlive(), c.neighboors.TotalAlive())
			}
		})
	}
}

func TestRule2CellLivingOn(t *testing.T) {
	cases := []TestCace{
		{true, Cells{{true}, {true}}, true},
		{true, Cells{{true}, {true}, {true}, {false}}, true},
	}

	for _, c := range cases {
		t.Run("Testing Rule 2", func(t *testing.T) {
			cell := NewCell(c.startAlive)
			newcell := cell.Interact(c.neighboors)
			if newcell.IsAlive() != c.endAlive {
				t.Errorf("Rule 2. Any live cell with two or three live neighbors lives on to the next generation. Cell Alive?: %v, Neighboors: %d", newcell.IsAlive(), c.neighboors.TotalAlive())
			}
		})
	}
}

func TestRule3CellOvercrowding(t *testing.T) {
	cell := NewCell(true)
	neighboors := Cells{{true}, {true}, {true}, {true}}
	newcell := cell.Interact(neighboors)
	if newcell.IsAlive() {
		t.Errorf("Rule 3. Any live cell with more than three live neighbors dies, as if by overcrowding.")
	}
}

func TestRule4CellReproduction(t *testing.T) {
	cases := []TestCace{
		{false, Cells{{true}, {true}}, false},
		{false, Cells{{true}, {true}, {true}}, true},
		{false, Cells{{true}, {true}, {true}, {true}}, false},
	}

	for _, c := range cases {
		t.Run("Testing Rule 4", func(t *testing.T) {
			cell := NewCell(c.startAlive)
			newcell := cell.Interact(c.neighboors)
			if newcell.IsAlive() != c.endAlive {
				t.Errorf("Rule 4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction. Cell Alive?: %v, Neighboors: %d", newcell.IsAlive(), c.neighboors.TotalAlive())
			}
		})
	}
}
