package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func execute(ctx context.Context, c *cli.Command) error {
	fmt.Println("Deleting a server!")
	return nil
}

var DeleteCommand = &cli.Command{
	Name:   "delete",
	Usage:  "Delete a server",
	Action: execute,
}
