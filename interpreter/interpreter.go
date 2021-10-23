package interpreter

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"vilmos/pixel"
	"vilmos/stack"

	"gopkg.in/ini.v1"
)

var (
	ErrorFileExtension   = errors.New("error: target image must be .png")
	ErrorOpenImage       = errors.New("error: unable to open specified image")
	ErrorRandomGenerator = errors.New("error: trying to generate a random number with n <= 0")
	ErrorDecodeImage     = errors.New("error: unable to decode specified image")
	ErrorOutOfBounds     = errors.New("error: out of bounds")
	ErrorInvalidHex      = errors.New("error: invalid hex format")
	ErrorLoadConfig      = errors.New("error: unable to load config file")
	ErrorCloseFile       = errors.New("error: unable to close the file")
	ErrorInputScanning   = errors.New("error: problems reading input")
)

var OPERATIONS = map[string]*pixel.Pixel{
	"INPUT_INT":    {R: 255, G: 255, B: 255}, //#ffffff -> INPUT INT
	"OUTPUT_INT":   {R: 0, G: 0, B: 0},       //#000000 -> OUTPUT INT
	"SUM":          {R: 0, G: 206, B: 209},   //#00ced1 -> SUM
	"SUB":          {R: 255, G: 165, B: 0},   //#ffa500 -> SUBTRACTION
	"DIV":          {R: 138, G: 43, B: 226},  //#8a2be2 -> DIVISION
	"MUL":          {R: 139, G: 0, B: 0},     //#8b0000 -> MULTIPLICATION
	"MOD":          {R: 255, G: 218, B: 185}, //#ffdab9 -> MODULUS
	"RND":          {R: 0, G: 128, B: 0},     //#008000 -> RANDOM
	"AND":          {R: 236, G: 243, B: 220}, //#ecf3dc -> AND
	"OR":           {R: 183, G: 198, B: 230}, //#b7c6e6 -> OR
	"XOR":          {R: 245, G: 227, B: 215}, //#f5e3d7 -> XOR
	"NAND":         {R: 225, G: 211, B: 239}, //#e1d3ef -> NAND
	"NOT":          {R: 255, G: 154, B: 162}, //#ff9aa2 -> NOT
	"INPUT_ASCII":  {R: 227, G: 227, B: 227}, //#e3e3e3 -> INPUT ASCII
	"OUTPUT_ASCII": {R: 75, G: 75, B: 75},    //#4b4b4b -> OUTPUT ASCII
	"POP":          {R: 204, G: 158, B: 6},   //#cc9e06 -> POP
	"SWAP":         {R: 255, G: 189, B: 74},  //#ffbd4a -> SWAP
	"CYCLE":        {R: 227, G: 127, B: 157}, //#e37f9d -> CYCLE
	"RCYCLE":       {R: 233, G: 148, B: 174}, //#e994ae -> RCYCLE
	"DUP":          {R: 0, G: 105, B: 148},   //#006994 -> DUPLICATE
	"REVERSE":      {R: 165, G: 165, B: 141}, //#a5a58d -> REVERSE
	"QUIT":         {R: 183, G: 228, B: 199}, //#b7e4c7 -> QUIT PROGRAM
	"OUTPUT":       {R: 155, G: 34, B: 66},   //#9B2242 -> OUTPUT ALL STACK
	"WHILE":        {R: 46, G: 26, B: 71},    //#2e1a47 -> START WHILE LOOP
	"WHILE_END":    {R: 104, G: 71, B: 141},  //#68478d -> END WHILE LOOP
}

type Interpreter struct {
	image   image.Image
	stack   *stack.Stack
	pc      image.Point
	width   int
	height  int
	isDebug bool
}

func NewInterpreter(debug bool, configs string) *Interpreter {
	rand.Seed(time.Now().UnixNano())

	interpreter := new(Interpreter)
	interpreter.image = nil
	interpreter.pc = image.Point{X: 0, Y: 0}
	interpreter.stack = stack.NewStack()
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	interpreter.isDebug = debug

	if configs != "" {
		err := loadConfigs(configs)
		if err != nil {
			logError(err)
		}
	}
	return interpreter
}

func (i *Interpreter) LoadImage(path string) error {
	fileExtension := filepath.Ext(path)

	if fileExtension != ".png" {
		return ErrorFileExtension
	}

	f, err := os.Open(path)
	if err != nil {
		return ErrorOpenImage
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logError(ErrorCloseFile)
		}
	}(f)
	img, _, err := image.Decode(f)

	if err != nil {
		return ErrorDecodeImage
	} else {
		i.image = img

		i.width, i.height = i.image.Bounds().Max.X, i.image.Bounds().Max.Y
		return nil
	}
}

func (i *Interpreter) GetImage() image.Image {
	return i.image
}

func (i *Interpreter) Run() {
	err := error(nil)
	stepCount := 0
	for err == nil {
		_, msg := i.Step()
		stepCount++
		if i.isDebug {
			debug(i, stepCount, msg)
			_, e := fmt.Scanf("\n")
			if e != nil {
				logError(ErrorInputScanning)
				return
			}
		}
		err = i.increasePC()
	}
}

func (i *Interpreter) Step() (bool, string) {
	px := i.readPixel()
	msg := processPixel(px, i)
	return true, msg
}

func (i *Interpreter) readPixel() *pixel.Pixel {
	return rgbaToPixel(i.image.At(i.pc.X, i.pc.Y).RGBA())
}

func rgbaToPixel(r uint32, g uint32, b uint32, _ uint32) *pixel.Pixel {
	return &pixel.Pixel{R: uint8(r / 257), G: uint8(g / 257), B: uint8(b / 257)}
}

func processPixel(pixel *pixel.Pixel, i *Interpreter) string {
	//goland:noinspection GrazieInspection
	switch pixel.String() {
	case OPERATIONS["INPUT_INT"].String(): //Gets value from input as number and pushes it to the stack
		var val int
		_, err := fmt.Scanf("%d\n", &val)
		if err != nil {
			logError(ErrorInputScanning)
			break
		}
		i.stack.Push(val)
		if i.isDebug {
			return "Pushed " + strconv.Itoa(val) + " into the stack"
		} else {
			return ""
		}
	case OPERATIONS["INPUT_ASCII"].String(): //Gets value from input as ASCII char and pushes it to the stack
		var val rune
		_, err := fmt.Scanf("%c\n", &val)
		if err != nil {
			logError(ErrorInputScanning)
			break
		}
		i.stack.Push(int(val))
	case OPERATIONS["OUTPUT_INT"].String(): //Pops the top of the stack and outputs it as number
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
		} else {
			fmt.Printf("%d", val)
		}
	case OPERATIONS["OUTPUT_ASCII"].String(): //Pops the top of the stack and outputs it as ASCII char
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
		} else {
			fmt.Printf("%c", val)
		}
	case OPERATIONS["SUM"].String(): //Pops two numbers, adds them and pushes the result in the stack
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
	case OPERATIONS["SUB"].String(): //Pops two numbers, subtracts them and pushes the result in the stack
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
	case OPERATIONS["DIV"].String(): //Pops two numbers, divides them and pushes the result in the stack
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
	case OPERATIONS["MUL"].String(): //Pops two numbers, multiplies them and pushes the result in the stack
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
	case OPERATIONS["MOD"].String(): //Pops two numbers, and pushes the result of the modulus in the stack
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
	case OPERATIONS["RND"].String(): //Pops one number, and pushes in the stack a random number between [0, n[ where n is the number popped
		n, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		if n <= 0 {
			logError(ErrorRandomGenerator)
			break
		}
		random := rand.Intn(n)
		i.stack.Push(random)
	case OPERATIONS["AND"].String(): //Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
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
	case OPERATIONS["OR"].String(): //Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
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
	case OPERATIONS["XOR"].String(): //Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
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
	case OPERATIONS["NAND"].String(): //Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
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
	case OPERATIONS["NOT"].String(): //Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		result := Btoi(!Itob(v1))
		i.stack.Push(result)
	case OPERATIONS["POP"].String(): //Pops one number, and discardes it
		_, err := i.stack.Pop()
		if err != nil {
			logError(err)
		}
	case OPERATIONS["SWAP"].String(): //Swaps the top two items in the stack
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
	case OPERATIONS["CYCLE"].String(): //Cycles clockwise the stack
		i.stack.Cycle()
	case OPERATIONS["RCYCLE"].String(): //Cycles anti-clockwise the stack
		i.stack.RCycle()
	case OPERATIONS["DUP"].String(): //Duplicates the top of the stack
		val, err := i.stack.Pop()
		if err != nil {
			logError(err)
			break
		}
		i.stack.Push(val)
		i.stack.Push(val)
	case OPERATIONS["REVERSE"].String(): //Reverses the content of the stack
		i.stack.Reverse()
	case OPERATIONS["QUIT"].String(): //Exits the program
		fmt.Printf("\n")
		os.Exit(1)
	case OPERATIONS["OUTPUT"].String(): //Outputs all the content of the stack without popping it
		i.stack.Output()
	case OPERATIONS["WHILE"].String():
		if i.stack.Peek() == 0 { //exits the loop if top is false
			jumpForward(i)
		}
	case OPERATIONS["WHILE_END"].String():
		jumpBack(i)
	default: //every color not in the list above pushes into the stack the sum of red, green and blue values of the pixel
		sum := pixel.R + pixel.G + pixel.B
		i.stack.Push(int(sum))
		return "Pushed " + strconv.Itoa(int(sum)) + " into the stack"
	}
	return ""
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
		case OPERATIONS["WHILE"].String():
			open++
		case OPERATIONS["WHILE_END"].String():
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
	closed := 0
	for {
		p := i.readPixel()
		err := i.decreasePC()
		switch p.String() {
		case OPERATIONS["WHILE"].String():
			closed--
			if closed == 0 {
				return
			}
		case OPERATIONS["WHILE_END"].String():
			closed++
		}
		if err != nil {
			logError(errors.New("error: missing start loop"))
		}
	}
}

func (i *Interpreter) increasePC() error {
	if i.pc.X+1 < i.width {
		i.pc.X = i.pc.X + 1
		return nil
	}
	if i.pc.Y+1 < i.height {
		i.pc.Y = i.pc.Y + 1
		i.pc.X = 0
		return nil
	}
	return ErrorOutOfBounds
}

func (i *Interpreter) decreasePC() error {
	if i.pc.X-1 >= 0 {
		i.pc.X = i.pc.X - 1
		return nil
	}
	if i.pc.Y-1 >= 0 {
		i.pc.Y = i.pc.Y - 1
		i.pc.X = i.width - 1
		return nil
	}
	return ErrorOutOfBounds
}

func hexToPixel(s string) (p *pixel.Pixel, err error) {
	var r, g, b int
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%2x%2x%2x", &r, &g, &b)
		if err != nil {
			return nil, ErrorInvalidHex
		}
		return &pixel.Pixel{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &r, &g, &b)
		if err != nil {
			return nil, ErrorInvalidHex
		}
		// Double the hex digits:
		r *= 17
		g *= 17
		b *= 17
		return &pixel.Pixel{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	default:
		err = ErrorInvalidHex
		return nil, err
	}
}

func loadConfigs(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return ErrorLoadConfig
	}
	for op := range OPERATIONS {
		value := cfg.Section("Colors").Key(op).String()
		if len(value) != 0 {
			newPx, err := hexToPixel(value)
			if err != nil {
				return ErrorInvalidHex
			}
			OPERATIONS[op] = newPx
		}
	}
	return err
}

func debug(i *Interpreter, step int, message string) {
	fmt.Printf("\n############ Step %d ############\n", step)
	fmt.Printf("Message: \033[33m%s\033[0m", message)
	for index := i.stack.Size() - 1; index >= 0; index-- {
		fmt.Printf("\n| %d  |", i.stack.GetItemAt(index))
	}
}
