package result

type Fireshot struct {
	grid       [7][7]string
	shotCoords []Coordinates
	hitShip    bool
}

type Coordinates struct {
	row int
	col int
}

func (f *Fireshot) shot(c Coordinates) {
	if f.grid[c.row][c.col] == "Ship" {
		f.hitShip = true
	}
}
