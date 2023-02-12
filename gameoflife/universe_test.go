package gameoflife

import (
	"testing"
)

func TestCanCreateARandomUniverse(t *testing.T) {
	width := 10
	height := 15
	u := NewRandomUniverse(width, height)
	if len(u) != height || len(u[0]) != width {
		t.Errorf("Expected universe to have a size of (%dx%d) but got (%dx%d)", width, height, len(u), len(u[0]))
	}
}

func TestCanCreateAGliderUniverse(t *testing.T) {
	size := 25
	u := NewGliderUniverse()
	if len(u) != size || len(u[0]) != size {
		t.Errorf("Expected universe to be (%dx%d) but got (%dx%d)", size, size, len(u), len(u[0]))
	}
	center := size / 2
	if !u.Get(center+1, center).IsAlive() || !u.Get(center+1, center-2).IsAlive() {
		t.Errorf("Could not find the glider pattern in the middle of the universe")
	}
}

func TestItCanCloneItself(t *testing.T) {
	u := NewUniverse(2, 2)
	u.Set(0, 0, true)

	u2 := u.Clone()
	if !u2.Get(0, 0).IsAlive() {
		t.Errorf("Cloned instance should contain the same data")
	}

	u2.Set(0, 0, false)
	if !u.Get(0, 0).IsAlive() {
		t.Errorf("Cloned instance changes should not affect the original")
	}
}

func TestItReturnsTheCorrectNeighboors(t *testing.T) {
	u := NewUniverse(3, 3)
	u.Set(0, 1, true)
	u.Set(1, 0, true)
	u.Set(2, 1, true)
	u.Set(1, 2, true)
	n := u.GetNeighboors(1, 1)
	if n.TotalAlive() != 4 {
		t.Errorf("Expected 4 live neightboors, got %d", n.TotalAlive())
	}

	n = u.GetNeighboors(0, 0)
	if n.TotalAlive() != 2 {
		t.Errorf("Expected 2 live neightboors, got %d", n.TotalAlive())
	}
}

func TestGliderPatternEvolution(t *testing.T) {
	size := 25
	u := NewGliderUniverse()

	center := size / 2
	if !u.Get(center+1, center).IsAlive() || !u.Get(center+1, center-2).IsAlive() {
		t.Errorf("Could not find the glider pattern in the middle of the universe")
	}

	u2 := u.Evolve()
	// check for deaths
	if u2.Get(center, center).IsAlive() {
		t.Error("Expected cell to die")
	}

	// check for no changes
	if !u2.Get(center+1, center).IsAlive() {
		t.Error("Expected cell to continue living")
	}

	// check for resurections
	if !u2.Get(center, center-1).IsAlive() {
		t.Error("Expected cell to become alive")
	}
}

func TestUniverseIsCorrectlyConvertedToString(t *testing.T) {
	u := NewUniverse(2, 2)
	u.Set(1, 1, true)
	s := u.ToString()

	expected := "  \n *\n"
	if expected != s {
		t.Errorf("String formatting does not match. Found: %s, Expected: %s", s, expected)
	}
}
