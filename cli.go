package main

import (
	"fmt"
	"log"
	"os"

	"./actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "varenv"
	app.Usage = "Create persistent env variables easily and quickly for testing locally."
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:  "load",
			Usage: "The yaml file where the variables live in, from `FILE`",
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					err := actions.Loader(c.Args().First())
					if err != nil {
						return err
					}
					return nil
				}
				log.Println("No argument given...")
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all the persistent environment variables on your computer.",
			Action: func(c *cli.Context) error {
				fmt.Println("\nList all env variables created by VARENV...\n")
				err := actions.List()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Quick commands to add an env variables through the cli interface",
			Action: func(c *cli.Context) error {
				if c.Args().Present() {
					fmt.Println("\nAdd env variables...\n")
					err := actions.Add(c.Args())
					if err != nil {
						return err
					}
					return nil
				}
				log.Println("\tNo argument given")
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Quick commands to remove an env variables through the cli interface",
			Action: func(c *cli.Context) error {
				fmt.Println("Remove env variables ", c.Args())
				actions.Remove()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
