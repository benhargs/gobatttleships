package result

type PlayerGrid struct {
	Grid   [7][7]string
	Coords Shot
}

type Shot struct {
	Row int
	Col int
}

/*
func (p *PlayerGrid) Shoot(s Shot) {
	if p.Grid[s.Row][s.Col] == "Ship" {
		p.HitShip = true
	}
}
*/
