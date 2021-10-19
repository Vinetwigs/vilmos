package main

import (
	"fmt"
	inter "vilmos/interpreter"
)

func main() {
	i := inter.NewInterpreter()

	err := i.LoadImage("C:\\Users\\User\\Desktop\\hello.jpg")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	i.Run()
}
