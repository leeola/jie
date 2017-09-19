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
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "stdin",
			Usage: "use stdin as the request body",
		},
		cli.BoolFlag{
			Name:  "stdout",
			Usage: "send response body to stdout",
		},
	}

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
		{
			Name:      "put",
			ArgsUsage: "ADDR",
			Usage:     "perform a PUT request",
			Action:    PutCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
