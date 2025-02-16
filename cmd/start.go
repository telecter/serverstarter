package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func start(ctx context.Context, c *cli.Command) error {
	fmt.Println("Starting a server...")
	return nil
}

var StartCommand = &cli.Command{
	Name:   "start",
	Usage:  "Start a server",
	Action: start,
}
