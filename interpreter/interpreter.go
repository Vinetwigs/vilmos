package interpreter

import (
	"image"
	"os"
	stack "vilmos/stack"
)

const (
	INT_TYPE = iota
	STRING_TYPE
)

type Interpreter struct {
	image       image.Image
	intstack    stack.IntStack
	stringstack stack.StringStack
	_type       int
}

func NewInterpreter() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.image = nil
	interpreter._type = INT_TYPE
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
		return nil
	}
}

func (i *Interpreter) GetImage() image.Image {
	return i.image
}
