package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func Input_str(Q string) string {
	fmt.Print(Q)
	// Scannerを使って一行読み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func file_exit(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {

	app := &cli.App{
		Name: "Renban-chan",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "izp",
				Aliases: []string{"z"},
				Value:   false,
				Usage:   "Index of zero padding",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "Interactive-mode",
				Aliases: []string{"i"},
				Usage:   "Interactive mode.",
				Action: func(c *cli.Context) error {
					root_path := ""
					if c.Args().Len() >= 1 {
						root_path = c.Args().Get(0)
					}
					if !file_exit(root_path) {
						os.Mkdir(root_path, 0777)
					}

					URL_template := Input_str("URL template(http://example.com/imgs/**.jpg):")

					S_index_s := Input_str("Start index:")
					S_index, err := strconv.Atoi(S_index_s)
					if err != nil {
						log.Fatalln(err)
					}

					E_index_s := Input_str("End index:")
					E_index, err := strconv.Atoi(E_index_s)
					if err != nil {
						log.Fatalln(err)
					}

					for i := S_index; i <= E_index; i++ {
						URL := Make_URL_by_index(URL_template, i)
						fmt.Printf("Downloading %s \n", URL)
						download_run(URL, i, root_path)
					}
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
