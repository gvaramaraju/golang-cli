package main

import (
	"fmt"
	"log"
	"os"

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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Usage: "Used to specify lang of output text",
				Value: "english",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
