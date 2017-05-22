package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jie"
	app.Usage = "a human friend json http client"
	app.Action = GetCmd

	app.Commands = []cli.Command{
		{
			Name:      "get",
			ArgsUsage: "ADDR",
			Aliases:   []string{"g"},
			Usage:     "get request",
			Action:    GetCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
