package main

import (
	"code"
	"context"
	"errors"
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
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "first_file",
				Value:     "",
				UsageText: "first file",
			},
			&cli.StringArg{
				Name:      "second_file",
				Value:     "",
				UsageText: "second file",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			firstFile := cmd.StringArg("first_file")
			secondFile := cmd.StringArg("second_file")

			if firstFile == "" || secondFile == "" {
				return errors.New("not enough arguments")
			}

			firstFileMap, err := code.ParseFile(firstFile)

			if err != nil {
				return err
			}

			secondFileMap, err := code.ParseFile(secondFile)

			if err != nil {
				return err
			}

			fmt.Println(firstFileMap, secondFileMap)

			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
