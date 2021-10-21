package pair

type Pair struct {
	x int
	y int
}

func NewPair(x int, y int) *Pair {
	p := new(Pair)
	p.x = x
	p.y = y
	return p
}

func (p *Pair) GetX() int {
	return p.x
}

func (p *Pair) GetY() int {
	return p.y
}
