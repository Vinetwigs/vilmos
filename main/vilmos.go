package main

import (
	//"flag"
	"errors"
	"fmt"
	"log"
	"os"
	inter "vilmos/interpreter"

	"github.com/urfave/cli/v2"
)

var (
	ErrorNoImage = errors.New("error: no specified image")
)

func main() {
	var debug bool

	app := &cli.App{
		Name:  "vilmos",
		Usage: "[WIP]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Value:       false,
				Usage:       "enable debug mode",
				Destination: &debug,
			},
		},
		Action: func(c *cli.Context) error {
			var file_path string
			if c.NArg() > 0 {
				file_path = c.Args().Get(0)
				i := inter.NewInterpreter(debug)

				err := i.LoadImage(file_path)
				if err != nil {
					logError(err)
					os.Exit(1)
				}
				i.Run()
			} else {
				logError(ErrorNoImage)
				os.Exit(1)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func logError(e error) {
	fmt.Printf("\n")
	log.Println("\033[31m" + e.Error() + "\033[0m")
	os.Exit(2)
}
