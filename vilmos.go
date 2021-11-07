package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	inter "github.com/Vinetwigs/vilmos/v2/interpreter"

	"github.com/urfave/cli/v2"
)

const (
	version string = "2.0.1"
	usage   string = "Official vilmos language interpreter"
)

var (
	ErrorNoImage      = errors.New("error: no specified image")
	ErrorWrongCommand = errors.New("error: wrong command")
)

func main() {
	var (
		debug           bool
		configPath      string
		maxSize         int
		instructionSize int
	)

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V", "v"},
		Usage:   "Shows installed version",
	}

	app := &cli.App{
		Name:    "vilmos",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Vinetwigs",
				Email: "github.com/Vinetwigs",
			},
		},
		Usage: usage,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Value:       false,
				Usage:       "enable debug mode",
				Destination: &debug,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"conf", "c"},
				Usage:       "load configuration from `FILE_PATH` for custom color codes",
				Value:       "",
				Destination: &configPath,
			},
			&cli.IntFlag{
				Name:        "max_size",
				Aliases:     []string{"m"},
				Usage:       "set max memory size",
				Value:       -1,
				Destination: &maxSize,
			},
			&cli.IntFlag{
				Name:        "instruction_size",
				Aliases:     []string{"is", "size", "s"},
				Usage:       "set instruction size",
				Value:       1,
				Destination: &instructionSize,
			},
		},
		Action: func(c *cli.Context) error {
			var filePath string
			if c.NArg() > 0 {
				filePath = c.Args().Get(0)
				i := inter.NewInterpreter(debug, configPath, maxSize, instructionSize)

				err := i.LoadImage(filePath)
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
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show installed version",
				Action: func(c *cli.Context) error {
					fmt.Printf("vilmos %s\n", version)
					return nil
				},
			},
		},
		CommandNotFound: func(c *cli.Context, command string) {
			logError(ErrorWrongCommand)
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
