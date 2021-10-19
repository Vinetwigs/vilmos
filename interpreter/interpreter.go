package interpreter

import (
	"image"
	"os"
)

type Interpreter struct {
	image image.Image
}

func NewInterpreter() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.image = nil
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
		return nil
	}
}

func (i *Interpreter) GetImage() image.Image {
	return i.image
}
