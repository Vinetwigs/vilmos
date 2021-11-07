package interpreter

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/ini.v1"
)

/*
 * Interpreter's throwable errors
 */
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
	ErrorOpenFile        = errors.New("error: unable to open specified file")
	ErrorWriteFile       = errors.New("error: error on writing on opened file")
	ErrorReadFile        = errors.New("error: error on reading opened file")
	ErrorInvalidString   = errors.New("error: invalid string into the stack")
	ErrorFileAlreadyOpen = errors.New("error: trying to open multiple files")
)

/*
 * A map of all interpreter's operations
 */
var OPERATIONS = map[string]*Pixel{
	"INPUT_INT":    {R: 255, G: 255, B: 255}, //#ffffff -> INPUT INT
	"OUTPUT_INT":   {R: 0, G: 0, B: 1},       //#000001 -> OUTPUT INT
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
	"FILE_OPEN":    {R: 145, G: 246, B: 139}, //#91f68b -> OPEN FILE
	"FILE_CLOSE":   {R: 47, G: 237, B: 35},   //#2fed23 -> CLOSE FILE
}

// Interpreter structure
type Interpreter struct {
	image           image.Image
	stack           *Stack
	pc              image.Point
	width           int
	height          int
	isDebug         bool
	instructionSize int
	openedFile      *os.File
}

// Interpreter's constructor. Params are flags value from CLI app.
func NewInterpreter(debug bool, configs string, maxSize int, instructionSize int) *Interpreter {
	rand.Seed(time.Now().UnixNano())

	interpreter := &Interpreter{
		image:           nil,
		stack:           NewStack(maxSize),
		pc:              image.Point{X: 0, Y: 0},
		width:           0,
		height:          0,
		isDebug:         debug,
		instructionSize: instructionSize,
		openedFile:      nil,
	}
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	if configs != "" {
		err := loadConfigs(configs)
		if err != nil {
			logError(err)
		}
	}
	return interpreter
}

// Loads image from OS and puts the stream into the interpreter image reference
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
	}
	i.image = img
	i.width, i.height = i.image.Bounds().Max.X, i.image.Bounds().Max.Y
	return nil
}

/*
 * Executes the image interpretation doing Step() while the image program is terminated.
 * It is responsible to increase the program counter and calling the debugger if the flag is set.
 */
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
			}
		}
		err = i.increasePC()
	}
}

// Interprets and executes next pixel in the given image
func (i *Interpreter) Step() (bool, string) {
	px := i.readPixel()
	msg := processPixel(px, i)
	return true, msg
}

// Reads pixel pointed by program counter and returns a Pixel struct reference
func (i *Interpreter) readPixel() *Pixel {
	return rgbaToPixel(i.image.At(i.pc.X, i.pc.Y).RGBA())
}

// Creates and returns a Pixel structure reference by rgba values
func rgbaToPixel(r uint32, g uint32, b uint32, _ uint32) *Pixel {
	return &Pixel{R: uint8(r / 257), G: uint8(g / 257), B: uint8(b / 257)}
}

// Tries to pop the stack. If it fails, an error will the throwed
func popOrErr(i *Interpreter) int {
	val, err := i.stack.Pop()
	if err != nil {
		logError(err)
	}
	return val
}

// Tries to push an item in the stack. If it fails, an error will be throwed.
func pushOrErr(i *Interpreter, val int) {
	err := i.stack.Push(val)
	if err != nil {
		logError(err)
	}
}

// Tries to read input from a given format. If it fails, an error will be throwed.
func scanfOrErr(format string, a *int) {
	_, err := fmt.Scanf(format, a)
	if err != nil {
		logError(ErrorInputScanning)
	}
}

// Executes a given pixel. Returns a message for the debugging.
func processPixel(pixel *Pixel, i *Interpreter) string {
	switch pixel.String() {
	case OPERATIONS["INPUT_INT"].String(): //Gets value from input as number and pushes it to the stack
		var val int
		if hasOpenedFile(i) {

			content, err := readFromFile(i)
			if err != nil {
				logError(err)
			}
			if i.isDebug {
				return "Pushed " + truncateString(content, 50) + " into the stack"
			}

		} else {
			scanfOrErr("%d\n", &val)
			pushOrErr(i, val)
		}
		if i.isDebug {
			return "Pushed " + strconv.Itoa(val) + " into the stack"
		}
	case OPERATIONS["INPUT_ASCII"].String(): //Gets values as ASCII char of a string and puts them into the stack
		var val string
		if hasOpenedFile(i) {

			content, err := readFromFile(i)
			if err != nil {
				logError(err)
			}
			if i.isDebug {
				return "Pushed " + truncateString(content, 50) + " into the stack"
			}

		} else {
			_, err := fmt.Scanf("%s\n", &val)
			if err != nil {
				logError(ErrorInputScanning)
			}

			pushOrErr(i, int('\000'))
			for _, char := range val {
				pushOrErr(i, int(char))
			}

		}

		if i.isDebug {
			return "Pushed " + val + " into the stack"
		}
	case OPERATIONS["OUTPUT_INT"].String(): //Pops the top of the stack and outputs it as number
		val := popOrErr(i)
		if hasOpenedFile(i) {
			_, err := i.openedFile.WriteString(strconv.Itoa(val))
			if err != nil {
				logError(ErrorWriteFile)
			}
			return "Wrote " + strconv.Itoa(val) + " to the opened file (" + i.openedFile.Name() + ")"
		} else {
			fmt.Printf("%d", val)
		}
		if i.isDebug {
			return "Popped " + strconv.Itoa(val) + " from the stack and printed it in the console"
		}
	case OPERATIONS["OUTPUT_ASCII"].String(): //Pops the top of the stack and outputs it as ASCII char
		str, err := buildStringFromStack(i)
		if err != nil {
			logError(ErrorInvalidString)
		}
		if hasOpenedFile(i) {
			_, err := i.openedFile.WriteString(str)
			if err != nil {
				logError(ErrorWriteFile)
			}
			return "Wrote " + str + " to the opened file (" + i.openedFile.Name() + ")"
		} else {
			fmt.Printf("%s", str)
		}
		if i.isDebug {
			return "Popped " + str + " from the stack and printed it in the console"
		}
	case OPERATIONS["SUM"].String(): //Pops two numbers, adds them and pushes the result in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		sum := v1 + v2
		pushOrErr(i, sum)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack their sum (" + strconv.Itoa(sum) + ")"
		}
	case OPERATIONS["SUB"].String(): //Pops two numbers, subtracts them and pushes the result in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		sub := v2 - v1
		pushOrErr(i, sub)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack their difference (" + strconv.Itoa(sub) + ")"
		}
	case OPERATIONS["DIV"].String(): //Pops two numbers, divides them and pushes the result in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		div := v2 / v1
		pushOrErr(i, div)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their division (" + strconv.Itoa(div) + ")"
		}
	case OPERATIONS["MUL"].String(): //Pops two numbers, multiplies them and pushes the result in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		mul := v1 * v2
		pushOrErr(i, mul)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack their multiplication (" + strconv.Itoa(mul) + ")"
		}
	case OPERATIONS["MOD"].String(): //Pops two numbers, and pushes the result of the modulus in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		mod := v2 % v1
		pushOrErr(i, mod)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their modulus (" + strconv.Itoa(mod) + ")"
		}
	case OPERATIONS["RND"].String(): //Pops one number, and pushes in the stack a random number between [0, n[ where n is the number popped
		n := popOrErr(i)
		if n <= 0 {
			logError(ErrorRandomGenerator)
		}
		random := rand.Intn(n)
		pushOrErr(i, random)
		if i.isDebug {
			return "Random generated " + strconv.Itoa(random) + " [range 0 to " + strconv.Itoa(n-1) + "] and then pushed it into the stack"
		}
	case OPERATIONS["AND"].String(): //Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		result := Itob(v1) && Itob(v2)
		pushOrErr(i, Btoi(result))
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their AND (" + strconv.Itoa(Btoi(result)) + ")"
		}
	case OPERATIONS["OR"].String(): //Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		result := Itob(v1) || Itob(v2)
		pushOrErr(i, Btoi(result))
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their OR (" + strconv.Itoa(Btoi(result)) + ")"
		}
	case OPERATIONS["XOR"].String(): //Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		result := Itob(v1) != Itob(v2)
		pushOrErr(i, Btoi(result))
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their XOR (" + strconv.Itoa(Btoi(result)) + ")"
		}
	case OPERATIONS["NAND"].String(): //Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		result := nand(Itob(v1), Itob(v2))
		pushOrErr(i, Btoi(result))
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and then pushed into the stack the result of their NAND (" + strconv.Itoa(Btoi(result)) + ")"
		}
	case OPERATIONS["NOT"].String(): //Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]
		v1 := popOrErr(i)
		result := Btoi(!Itob(v1))
		pushOrErr(i, result)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + " from the stack and then pushed into the stack its NOT (" + strconv.Itoa(result) + ")"
		}
	case OPERATIONS["POP"].String(): //Pops one number, and discardes it
		v := popOrErr(i)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v) + " from the stack"
		}
	case OPERATIONS["SWAP"].String(): //Swaps the top two items in the stack
		v1 := popOrErr(i)
		v2 := popOrErr(i)
		pushOrErr(i, v1)
		pushOrErr(i, v2)
		if i.isDebug {
			return "Popped " + strconv.Itoa(v1) + ", popped " + strconv.Itoa(v2) + " and pushed in reverse order to swap them"
		}
	case OPERATIONS["CYCLE"].String(): //Cycles clockwise the stack
		i.stack.Cycle()
		if i.isDebug {
			return "Cycled clockwise by one step the stack"
		}
	case OPERATIONS["RCYCLE"].String(): //Cycles anti-clockwise the stack
		i.stack.RCycle()
		if i.isDebug {
			return "Cycled counter-clockwise by one step the stack"
		}
	case OPERATIONS["DUP"].String(): //Duplicates the top of the stack
		val := popOrErr(i)
		pushOrErr(i, val)
		pushOrErr(i, val)
		if i.isDebug {
			return "Popped " + strconv.Itoa(val) + " and then pushed it twice to duplicate it"
		}
	case OPERATIONS["REVERSE"].String(): //Reverses the content of the stack
		i.stack.Reverse()
		if i.isDebug {
			return "Reversed stack content"
		}
	case OPERATIONS["QUIT"].String(): //Exits the program
		fmt.Printf("\n")
		os.Exit(0)
	case OPERATIONS["OUTPUT"].String(): //Outputs all the content of the stack without popping it
		i.stack.Output()
		if i.isDebug {
			return "Outputted all the stack content"
		}
	case OPERATIONS["WHILE"].String():
		if i.stack.Peek() == 0 { //exits the loop if top is false
			jumpForward(i)
			if i.isDebug {
				return "Jumped forward for while loop"
			}
		}
		if i.isDebug {
			return "Entered in while loop"
		}
	case OPERATIONS["WHILE_END"].String():
		jumpBack(i)
		if i.isDebug {
			return "Jumped back for while loop"
		}
	case OPERATIONS["FILE_OPEN"].String():
		if hasOpenedFile(i) {
			logError(ErrorFileAlreadyOpen)
		}
		var err error = nil
		fileName, err := buildStringFromStack(i)
		if err != nil {
			logError(err)
		}
		i.openedFile, err = openFile(fileName)
		if err != nil {
			logError(err)
		}
		return "Opened file " + i.openedFile.Name()
	case OPERATIONS["FILE_CLOSE"].String():
		fileName := i.openedFile.Name()
		err := i.openedFile.Close()
		if err != nil {
			logError(ErrorCloseFile)
		}
		i.openedFile = nil
		return "Closed file " + fileName
	default: //every color not in the list above pushes into the stack the sum of red, green and blue values of the pixel
		sum := pixel.R + pixel.G + pixel.B
		pushOrErr(i, int(sum))
		return "Pushed " + strconv.Itoa(int(sum)) + " into the stack"
	}
	return ""
}

// Logs an error to stdout and stops the interpreter
func logError(e error) {
	fmt.Printf("\n")
	log.Println("\033[31m" + e.Error() + "\033[0m")
	os.Exit(2)
}

// Converts an integer to bool
func Itob(i int) bool {
	return i != 0
}

// Converts a bool to integer
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Returns the result of a NAND b
func nand(a bool, b bool) bool {
	return !(a && b)
}

// Jumps forward to the corresponding end while operation
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

// Jumps back to the corresponding open while operation
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

// Increases program counter
func (i *Interpreter) increasePC() error {
	if i.pc.X+i.instructionSize < i.width {
		i.pc.X = i.pc.X + i.instructionSize
		return nil
	}
	if i.pc.Y+i.instructionSize < i.height {
		i.pc.Y = i.pc.Y + i.instructionSize
		i.pc.X = 0
		return nil
	}
	return ErrorOutOfBounds
}

// Decreases program counter
func (i *Interpreter) decreasePC() error {
	if i.pc.X-i.instructionSize >= 0 {
		i.pc.X = i.pc.X - i.instructionSize
		return nil
	}
	if i.pc.Y-i.instructionSize >= 0 {
		i.pc.Y = i.pc.Y - i.instructionSize
		i.pc.X = i.width - i.instructionSize
		return nil
	}
	return ErrorOutOfBounds
}

// Converts a string representing an hex value to a Pixel structure. An error will be throwed if the format is wrong.
func hexToPixel(s string) (p *Pixel, err error) {
	var r, g, b int
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%2x%2x%2x", &r, &g, &b)
		if err != nil {
			return nil, ErrorInvalidHex
		}
		return &Pixel{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &r, &g, &b)
		if err != nil {
			return nil, ErrorInvalidHex
		}
		// Double the hex digits:
		r *= 17
		g *= 17
		b *= 17
		return &Pixel{R: uint8(r), G: uint8(g), B: uint8(b)}, nil
	default:
		err = ErrorInvalidHex
		return nil, err
	}
}

// Loads configs from the given config file and overrides standard operations color codes with the custom ones
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

// Displays a debug message and the stack content in the specified step
func debug(i *Interpreter, step int, message string) {
	fmt.Printf("\n############ Step %d ############\n", step)
	fmt.Printf("Message: \033[33m%s\033[0m", message)
	for index := i.stack.Size() - 1; index >= 0; index-- {
		fmt.Printf("\n|%8d|", i.stack.GetItemAt(index))
	}
	fmt.Print("\nPress ENTER to step over:")
}

func buildStringFromStack(i *Interpreter) (string, error) {
	var (
		result string = ""
		ch     rune   = ' '
	)
	for index := i.stack.Size() - 1; index >= 0; index-- {
		ch = rune(popOrErr(i))
		if ch == '\000' {
			return result, nil
		}
		result += string(ch)
	}
	return "", ErrorInvalidString
}

func openFile(path string) (*os.File, error) {
	var (
		err  error
		file *os.File
	)
	x := os.O_CREATE | os.O_APPEND | os.O_RDWR
	file, err = os.OpenFile(path, x, 0644)
	if err != nil {
		return nil, ErrorOpenFile
	}
	return file, nil
}

func hasOpenedFile(i *Interpreter) bool {
	return i.openedFile != nil
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}

func readFromFile(i *Interpreter) (string, error) {
	reader := bufio.NewReader(i.openedFile)
	var ch rune = '1'

	var content string

	pushOrErr(i, int('\000'))
	for ch != '\000' {
		ch, _, err := reader.ReadRune()

		if err != nil {

			if err != io.EOF {

				return "", ErrorReadFile
			}

			break
		}

		content += string(ch)
		pushOrErr(i, int(ch))
	}
	return content, nil
}
