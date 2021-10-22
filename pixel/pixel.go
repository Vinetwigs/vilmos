package pixel

import "strconv"

type Pixel struct {
	R uint8
	G uint8
	B uint8
}

func (p *Pixel) String() string {
	return "{R: " + strconv.Itoa(int(p.R)) + " G: " + strconv.Itoa(int(p.G)) + " B: " + strconv.Itoa(int(p.B)) + "}"
}

func (p *Pixel) Equals(other Pixel) bool {
	return (p.R == other.R) && (p.G == other.G) && (p.B == other.B)
}
