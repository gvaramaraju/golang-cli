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
			fmt.Println("Printing CLI args")
			for _, arg := range c.Args().Slice() {
				fmt.Println("ARG ::", arg)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
