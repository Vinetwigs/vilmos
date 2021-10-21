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
	pair "vilmos/pair"
	pixel "vilmos/pixel"
	stack "vilmos/stack"
)

var (
	WHITE            pixel.Pixel = pixel.Pixel{R: 255, G: 255, B: 255, A: 255} //#ffffff -> INPUT INT
	BLACK            pixel.Pixel = pixel.Pixel{R: 0, G: 0, B: 0, A: 255}       //#000000 -> OUTPUT INT
	TURQUOISE        pixel.Pixel = pixel.Pixel{R: 0, G: 206, B: 209, A: 255}   //#00ced1 -> SUM
	ORANGE           pixel.Pixel = pixel.Pixel{R: 255, G: 165, B: 0, A: 255}   //#ffa500 -> SUBTRACTION
	VIOLET           pixel.Pixel = pixel.Pixel{R: 138, G: 43, B: 226, A: 255}  //#8a2be2 -> DIVISION
	RED              pixel.Pixel = pixel.Pixel{R: 139, G: 0, B: 0, A: 255}     //#8b0000 -> MULTIPLICATION
	PEACH            pixel.Pixel = pixel.Pixel{R: 255, G: 218, B: 185, A: 255} //#ffdab9 -> MODULUS
	GREEN            pixel.Pixel = pixel.Pixel{R: 0, G: 128, B: 0, A: 255}     //#008000 -> RANDOM
	BEIGE            pixel.Pixel = pixel.Pixel{R: 236, G: 243, B: 220, A: 255} //#ecf3dc -> AND
	LIGHT_STEEL_BLUE pixel.Pixel = pixel.Pixel{R: 183, G: 198, B: 230, A: 255} //#b7c6e6 -> OR
	WHITE_CHOCOLATE  pixel.Pixel = pixel.Pixel{R: 245, G: 227, B: 215, A: 255} //#f5e3d7 -> XOR
	PALE_LAVANDER    pixel.Pixel = pixel.Pixel{R: 225, G: 211, B: 239, A: 255} //#e1d3ef -> NAND
	SALMON           pixel.Pixel = pixel.Pixel{R: 255, G: 154, B: 162, A: 255} //#ff9aa2 -> NOT
	DARK_WHITE       pixel.Pixel = pixel.Pixel{R: 227, G: 227, B: 227, A: 255} //#e3e3e3 -> INPUT ASCII
	LIGHT_BLACK      pixel.Pixel = pixel.Pixel{R: 75, G: 75, B: 75, A: 255}    //#4b4b4b -> OUTPUT ASCII
	DARK_GOLD        pixel.Pixel = pixel.Pixel{R: 204, G: 158, B: 6, A: 255}   //#cc9e06 -> POP
	GOLD             pixel.Pixel = pixel.Pixel{R: 255, G: 189, B: 74, A: 255}  //#ffbd4a -> SWAP
	SHIMMERING_BLUSH pixel.Pixel = pixel.Pixel{R: 227, G: 127, B: 157, A: 255} //#e37f9d -> CYCLE
	CHARM_PINK       pixel.Pixel = pixel.Pixel{R: 233, G: 148, B: 174, A: 255} //#e994ae -> RCYCLE
	SEA_BLUE         pixel.Pixel = pixel.Pixel{R: 0, G: 105, B: 148, A: 255}   //#006994 -> DUPLICATE
	CAR_OIL          pixel.Pixel = pixel.Pixel{R: 165, G: 165, B: 141, A: 255} //#a5a58d -> REVERSE
	HAWAII_SEA       pixel.Pixel = pixel.Pixel{R: 183, G: 228, B: 199, A: 255} //#b7e4c7 -> QUIT PROGRAM
	WINE_RED         pixel.Pixel = pixel.Pixel{R: 155, G: 34, B: 66, A: 255}   //#9B2242 -> OUTPUT ALL STACK
	MIDNIGHT_PURPLE  pixel.Pixel = pixel.Pixel{R: 46, G: 26, B: 71, A: 255}    //#2e1a47 -> START WHILE LOOP
	ROYALE_PURPLE    pixel.Pixel = pixel.Pixel{R: 104, G: 71, B: 141, A: 255}  //#68478d -> END WHILE LOOP
)

var (
	WIDTH  int = 0
	HEIGTH int = 0
)

type Interpreter struct {
	image image.Image
	stack stack.Stack
	pc    pair.Pair
}

func NewInterpreter() *Interpreter {
	rand.Seed(time.Now().UnixNano())

	interpreter := new(Interpreter)
	interpreter.image = nil
	interpreter.pc = *pair.NewPair(0, 0)

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	return interpreter
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
	var err error = nil
	for err == nil {
		i.Step()
		err = i.increasePC()
	}
}

func (i *Interpreter) Step() bool {
	pixel := i.readPixel()
	processPixel(pixel, i)
	return true
}

func (i *Interpreter) readPixel() *pixel.Pixel {
	return rgbaToPixel(i.image.At(i.pc.GetX(), i.pc.GetY()).RGBA())
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) *pixel.Pixel {
	return &pixel.Pixel{R: int(r / 257), G: int(g / 257), B: int(b / 257), A: int(a / 257)}
}

func processPixel(pixel *pixel.Pixel, i *Interpreter) {
	switch pixel.String() {
	case WHITE.String(): //Gets value from input as number and pushes it to the stack
		var val int
		fmt.Scanf("%d\n", &val)
		i.stack.Push(val)
	case DARK_WHITE.String(): //Gets value from input as ASCII char and pushes it to the stack
		var val rune
		fmt.Scanf("%c\n", &val)
		i.stack.Push(int(val))
	case BLACK.String(): //Pops the top of the stack and outputs it as number
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
		} else {
			fmt.Printf("%d", val)
		}
	case LIGHT_BLACK.String(): //Pops the top of the stack and outputs it as ASCII char
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
		} else {
			fmt.Printf("%c", val)
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
		sub := v2 - v1
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
		sub := v2 / v1
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
		sub := v2 % v1
		i.stack.Push(sub)
	case GREEN.String(): //Pops one number, and pushes in the stack a random number between [0, n) where n is the number popped
		n, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		if n <= 0 {
			logError(errors.New("error: cannot generate a random number with n <= 0"))
			break
		}
		random := rand.Intn(n)
		i.stack.Push(random)
	case BEIGE.String(): //Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := Itob(v1) && Itob(v2)
		i.stack.Push(Btoi(result))
	case LIGHT_STEEL_BLUE.String(): //Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := Itob(v1) || Itob(v2)
		i.stack.Push(Btoi(result))
	case WHITE_CHOCOLATE.String(): //Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := Itob(v1) != Itob(v2)
		i.stack.Push(Btoi(result))
	case PALE_LAVANDER.String(): //Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := nand(Itob(v1), Itob(v2))
		i.stack.Push(Btoi(result))
	case SALMON.String(): //Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := Btoi(!Itob(v1))
		i.stack.Push(result)
	case DARK_GOLD.String(): //Pops one number, and discardes it
		_, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
	case GOLD.String(): //Swaps the top two items in the stack
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		v2, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
		i.stack.Push(v1)
		i.stack.Push(v2)
	case SHIMMERING_BLUSH.String(): //Cycles clockwise the stack
		i.stack.Cycle()
	case CHARM_PINK.String(): //Cycles anti-clockwise the stack
		i.stack.RCycle()
	case SEA_BLUE.String(): //Duplicates the top of the stack
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		i.stack.Push(val)
		i.stack.Push(val)
	case CAR_OIL.String(): //Reverses the content of the stack
		i.stack.Reverse()
	case HAWAII_SEA.String(): //Exits the program
		fmt.Printf("\n")
		os.Exit(1)
	case WINE_RED.String(): //Outputs all the content of the stack without popping it
		i.stack.Output()
	case MIDNIGHT_PURPLE.String():
		if i.stack.Peek() == 0 { //exits the loop if top is false
			jumpForward(i)
		}
	case ROYALE_PURPLE.String():
		jumpBack(i)
	default: //every color not in the list above pushes into the stack the sum of red, green and blue values of the pixel
		sum := pixel.R + pixel.G + pixel.B
		i.stack.Push(sum)
	}
}

func logError(e error) {
	fmt.Printf("\n")
	log.Println("\033[31m" + e.Error() + "\033[0m")
	os.Exit(2)
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

func nand(a bool, b bool) bool {
	return !(a && b)
}

func jumpForward(i *Interpreter) {
	open := 0
	for {
		p := i.readPixel()
		err := i.increasePC()
		switch p.String() {
		case MIDNIGHT_PURPLE.String():
			open++
		case ROYALE_PURPLE.String():
			open--
			if open == 0 {
				return
			}
		}
		if err != nil {
			logError(errors.New("error: missing end loop"))
		}
	}
}

func jumpBack(i *Interpreter) {
	close := 0
	for {
		p := i.readPixel()
		err := i.decreasePC()
		switch p.String() {
		case MIDNIGHT_PURPLE.String():
			close--
			if close == 0 {
				return
			}
		case ROYALE_PURPLE.String():
			close++
		}
		if err != nil {
			logError(errors.New("error: missing start loop"))
		}
	}
}

func (i *Interpreter) increasePC() error {
	if i.pc.GetX()+1 < WIDTH {
		i.pc.SetX(i.pc.GetX() + 1)
		return nil
	}
	if i.pc.GetY()+1 < HEIGTH {
		i.pc.SetY(i.pc.GetY() + 1)
		i.pc.SetX(0)
		return nil
	}
	return errors.New("error: out of bounds")
}

func (i *Interpreter) decreasePC() error {
	if i.pc.GetX()-1 >= 0 {
		i.pc.SetX(i.pc.GetX() - 1)
		return nil
	}
	if i.pc.GetY()-1 >= 0 {
		i.pc.SetY(i.pc.GetY() - 1)
		i.pc.SetX(WIDTH - 1)
		return nil
	}
	return errors.New("error: out of bounds")
}
