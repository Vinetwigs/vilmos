package interpreter

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
	pixel "vilmos/pixel"
	stack "vilmos/stack"
)

var (
	WHITE     pixel.Pixel = pixel.Pixel{R: 255, G: 255, B: 255, A: 255} //#ffffff -> INPUT INT
	BLACK     pixel.Pixel = pixel.Pixel{R: 0, G: 0, B: 0, A: 255}       //#000000 -> OUTPUT INT
	TURQUOISE pixel.Pixel = pixel.Pixel{R: 0, G: 206, B: 209, A: 255}   //#00ced1 -> SUM
	ORANGE    pixel.Pixel = pixel.Pixel{R: 255, G: 165, B: 0, A: 255}   //#ffa500 -> SUBTRACTION
	VIOLET    pixel.Pixel = pixel.Pixel{R: 138, G: 43, B: 226, A: 255}  //#8a2be2 -> DIVISION
	RED       pixel.Pixel = pixel.Pixel{R: 139, G: 0, B: 0, A: 255}     //#8b0000 -> MULTIPLICATION
	PEACH     pixel.Pixel = pixel.Pixel{R: 255, G: 218, B: 185, A: 255} //#ffdab9 -> MODULUS
	GREEN     pixel.Pixel = pixel.Pixel{R: 0, G: 128, B: 0, A: 255}     //#008000 -> RANDOM
	BEIGE     pixel.Pixel = pixel.Pixel{R: 236, G: 243, B: 220, A: 255} //#ecf3dc -> OR
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
	rand.Seed(time.Now().UnixNano())

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
		var val int
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
	case ORANGE.String(): //Pops two numbers, subtracts them and pushes the result in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		sub := v1 - v2
		i.stack.Push(sub)
	case VIOLET.String(): //Pops two numbers, divides them and pushes the result in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		sub := v1 / v2
		i.stack.Push(sub)
	case RED.String(): //Pops two numbers, multiplies them and pushes the result in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		sub := v1 * v2
		i.stack.Push(sub)
	case PEACH.String(): //Pops two numbers, and pushes the result of the modulus in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		sub := v1 % v2
		i.stack.Push(sub)
	case GREEN.String(): //Pops one number, and pushes in the stack a random number between [0, n) where n is the number popped
		n, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		if n <= 0 {
			logError(errors.New("error: cannot generate a random number with n <= 0"))
			break
		}
		random := rand.Intn(n)
		i.stack.Push(random)
	case BEIGE.String(): //Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		result := Itob(v1) && Itob(v2)
		i.stack.Push(Btoi(result))
	default: //every color not in the list above pushes into the stack the sum of red, green and blue values of the pixel
		sum := pixel.R + pixel.G + pixel.B
		i.stack.Push(sum)
	}
}

func logError(e error) {
	log.Println("\033[31m" + e.Error() + "\033[0m")
}

func Itob(i int) bool {
	return i != 0
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
