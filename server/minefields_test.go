package entrapped

import (
	"testing"
)

func TestCreateEmptyMineField(t *testing.T) {
	field := createEmptyMineField(9)

	if field.size != 9 {
		t.Error("wrong mine field size")
	}
}

func TestAddRandomBombs(t *testing.T) {
	setField := createEmptyMineField(8).addRandomBombs(9)
	t.Log(setField)
}
