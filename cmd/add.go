package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func add(ctx context.Context, c *cli.Command) error {
	if c.Args().Len() < 2 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}
	fmt.Printf("Creating a %s server with %s", "ballz", "cock")
	return nil
}

var AddCommand = &cli.Command{
	Name:      "add",
	Usage:     "Add a new server",
	ArgsUsage: "<version> <type>",
	Action:    add,
}
