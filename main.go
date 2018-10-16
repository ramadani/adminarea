package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "adminarea"
	app.Usage = "Administrative Areas"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "init",
			Usage: "initialization adminarea to create areas table",
			Action: func(c *cli.Context) error {
				fmt.Println("Setup")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
