package entrapped

type trap struct {
	trapMap *mineField
	lifes   int
}

func makeTrap(size, numBombs, lifes int) *trap {
	return &trap{
		trapMap: createEmptyMineField(size).addRandomBombs(numBombs),
		lifes:   lifes,
	}
}

func (t *trap) open(idx int) (int, int, string) {
	if idx > (t.trapMap.size*t.trapMap.size)-1 {
		return 0, 0, "error:invalid idx"
	}

	if t.lifes > 0 {
		ele := t.trapMap.checkIndex(idx)
		if ele == mine {
			t.lifes--
			return ele, t.lifes, ""
		} else {
			return ele, t.lifes, ""
		}
	}

	return 0, 0, ""
}
