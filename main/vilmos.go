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

const (
	version string = "1.0.0"
)

var (
	ErrorNoImage      = errors.New("error: no specified image")
	ErrorWrongCommand = errors.New("error: wrong command")
)

func main() {
	var debug bool
	var config_path string

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V", "v"},
		Usage:   "shows installed version",
	}

	app := &cli.App{
		Name:    "vilmos",
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Vinetwigs",
				Email: "github.com/Vinetwigs",
			},
		},
		Usage: "[WIP]",
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
				Usage:       "set config file path for custom color codes",
				Value:       "",
				DefaultText: "not set",
				Destination: &config_path,
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
