package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Golang CLI demo using urfave cli",
		Usage: "Demonstrate cli",
		Action: func(c *cli.Context) error {
			fmt.Println("Golang CLI ran successfully")
			lang := c.String("lang")
			fmt.Println("Printing CLI args")
			for _, arg := range c.Args().Slice() {
				fmt.Println("ARG ::", arg)
			}
			if lang == "spanish" {
				fmt.Println("Hola ! Have a great day!!")
			} else {
				fmt.Println("Hello ! Have a great day !!")
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "time",
				Usage: "Print current system time",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "lang",
						Aliases: []string{"l"},
						Usage:   "Used to specify `language` of output text",
						Value:   "english",
						// EnvVars:  []string{"LANG"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					lang := c.String("lang")
					if lang == "spanish" {
						fmt.Println("Hola User! Time is ", time.Now().String())
					} else {
						fmt.Println("Hello User! Time is ", time.Now().String())
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
