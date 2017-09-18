package main

import (
	"bytes"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/leeola/fixity/util/clijson"
	"github.com/urfave/cli"
)

func PostCmd(ctx *cli.Context) error {
	urlStr := ctx.Args().First()
	if urlStr == "" {
		return errors.New("missing url argument")
	}

	if !strings.HasPrefix(strings.ToLower(urlStr), "http") {
		urlStr = "http://" + urlStr
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	jsonArgs := ctx.Args().Tail()

	jsonB, err := clijson.CliJson(jsonArgs)
	if err != nil {
		return err
	}

	reqConf := Config{
		PipeResponse: ctx.GlobalBool("pipe-response"),
		Method:       "GET",
		URL:          u.String(),
		Body:         bytes.NewReader(jsonB),
		Writer:       os.Stdout,
	}
	return Request(reqConf)
}
