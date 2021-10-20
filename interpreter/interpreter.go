package interpreter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	pixel "vilmos/pixel"
	stack "vilmos/stack"
)

var (
	WHITE     pixel.Pixel = pixel.Pixel{R: 255, G: 255, B: 255, A: 255} //#ffffff
	BLACK     pixel.Pixel = pixel.Pixel{R: 0, G: 0, B: 0, A: 255}       //#000000
	TURQUOISE pixel.Pixel = pixel.Pixel{R: 0, G: 206, B: 209, A: 255}   //#00ced1
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
	image image.Image
	stack stack.Stack
	_type int
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
	processPixel(&pixel, i)
	return true
}

func readPixel(i *Interpreter, x int, y int) pixel.Pixel {
	return rgbaToPixel(i.image.At(x, y).RGBA())
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel.Pixel {
	return pixel.Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257), A: int(a / 257)}
}

func processPixel(pixel *pixel.Pixel, i *Interpreter) {
	switch pixel.String() {
	case WHITE.String(): //Gets value from input as number and pushes it to the stack
		var val uint8
		fmt.Scanf("%d\n", &val)
		i.stack.Push(val)
	case BLACK.String(): //Pops the top of the stack and outputs it as number
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
		} else {
			fmt.Printf("%d", val)
		}
	case TURQUOISE.String(): //Pops two numbers, adds them and pushes the result in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}

		sum := v1 + v2
		i.stack.Push(sum)
	default: //every color not in the list above pushes into the stack the sum of red, green and blue values of the pixel
		sum := pixel.R + pixel.G + pixel.B
		i.stack.Push(uint8(sum))
	}
}

func logError(e error) {
	log.Println("\033[31m" + e.Error() + "\033[0m")
}
