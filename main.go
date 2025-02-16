package main

import (
	"context"
	"log"
	"os"
	"serverstarter/cmd"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:        "serverstarter",
		Description: "Start various Minecraft servers with ease",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir",
				Usage: "Directory to use for servers",
			},
		},
		Commands: []*cli.Command{
			cmd.StartCommand,
			cmd.AddCommand,
			cmd.DeleteCommand,
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
