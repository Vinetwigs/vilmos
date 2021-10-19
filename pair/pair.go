package pair

type Pair struct {
	x int
	y int
}

func NewPair() *Pair {
	p := new(Pair)
	p.x = 0
	p.y = 0
	return p
}

func (p *Pair) SetX(val int) {
	p.x = val
}

func (p *Pair) SetY(val int) {
	p.y = val
}

func (p *Pair) GetX() int {
	return p.x
}

func (p *Pair) GetY() int {
	return p.y
}
