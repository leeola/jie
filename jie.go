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
			Usage:     "perform a GET request",
			Action:    GetCmd,
		},
		{
			Name:      "post",
			ArgsUsage: "ADDR",
			Usage:     "perform a POST request",
			Action:    PostCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
