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
	version string = "2.1.1"
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
		imagePath       string
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
				Usage:       "set max memory `SIZE`",
				Value:       -1,
				Destination: &maxSize,
			},
			&cli.IntFlag{
				Name:        "instruction_size",
				Aliases:     []string{"is", "size", "s"},
				Usage:       "set instruction `SIZE`",
				Value:       1,
				Destination: &instructionSize,
			},
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "input `FILE_PATH`",
				Value:       "",
				Destination: &imagePath,
			},
		},
		Action: func(c *cli.Context) error {
			if imagePath != "" {
				i := inter.NewInterpreter(debug, maxSize, instructionSize)

				if configPath != "" {
					err := inter.LoadConfigs(configPath)
					if err != nil {
						logError(err)
					}
				}

				err := i.LoadImage(imagePath)
				if err != nil {
					logError(err)
				}
				i.Run()
			} else {
				logError(ErrorNoImage)
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
