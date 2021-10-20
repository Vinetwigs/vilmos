package interpreter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	pixel "vilmos/pixel"
	stack "vilmos/stack"
)

var (
	WHITE pixel.Pixel = pixel.Pixel{R: 255, G: 255, B: 255, A: 255}
)

const (
	INT_TYPE = iota
	STRING_TYPE
)

var (
	WIDTH  int = 0
	HEIGTH int = 0
)

type Interpreter struct {
	image    image.Image
	intstack stack.Stack
	_type    int
}

func NewInterpreter() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.image = nil
	interpreter._type = INT_TYPE

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	return interpreter
}

/*
	0 = INT_TYPE
	> 0 = STRING_TYPE
*/
func (i *Interpreter) SetType(t int) {
	if t != 0 {
		i._type = INT_TYPE
	} else {
		i._type = STRING_TYPE
	}
}

func (i *Interpreter) LoadImage(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	if err != nil {
		return err
	} else {
		i.image = image

		WIDTH, HEIGTH = i.image.Bounds().Max.X, i.image.Bounds().Max.Y
		return nil
	}
}

func (i *Interpreter) GetImage() image.Image {
	return i.image
}

func (i *Interpreter) Run() {
	for y := 0; y < HEIGTH; y++ {
		for x := 0; x < WIDTH; x++ {
			i.Step(x, y)
		}
	}
}

func (i *Interpreter) Step(x int, y int) bool {
	pixel := readPixel(i, x, y)
	processPixel(&pixel)
	return true
}

func readPixel(i *Interpreter, x int, y int) pixel.Pixel {
	return rgbaToPixel(i.image.At(x, y).RGBA())
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel.Pixel {
	return pixel.Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257), A: int(a / 257)}
}

func processPixel(pixel *pixel.Pixel) {
	switch pixel.String() {
	case WHITE.String():
		fmt.Println("BIANCO")
	}
}
