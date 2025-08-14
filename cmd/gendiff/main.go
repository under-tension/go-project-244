package main

import (
	"context"
	"fmt"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "format",
			Usage:       "output format",
			Aliases:     []string{"f"},
			DefaultText: "stylish",
		},
	}
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: flags,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
