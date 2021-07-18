package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "Renban-chan",
		Commands: []*cli.Command{
			{
				Name:    "Interactive-mode",
				Aliases: []string{"i"},
				Usage:   "Interactive mode.",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {

					return nil
				},
			},
			{
				Name:    "Not-interactive-mode",
				Aliases: []string{"r"},
				Usage:   "Not interactive mode.",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {

					return nil
				},
			},
		},
	}

	app.Run(os.Args)

}
