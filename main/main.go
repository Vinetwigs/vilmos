package main

import (
	"fmt"
	inter "vilmos/interpreter"
)

func main() {
	i := inter.NewInterpreter()

	err := i.LoadImage("C:\\Users\\User\\Desktop\\sum_from_image.png")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	i.Run()
}
