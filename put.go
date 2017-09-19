package main

import (
	"bytes"
	"errors"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/leeola/fixity/util/clijson"
	"github.com/urfave/cli"
)

func PutCmd(ctx *cli.Context) error {
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

	var r io.Reader
	if !ctx.GlobalBool("stdin") {
		jsonArgs := ctx.Args().Tail()
		jsonB, err := clijson.CliJson(jsonArgs)
		if err != nil {
			return err
		}
		r = bytes.NewReader(jsonB)
	} else {
		r = os.Stdin
	}

	reqConf := Config{
		PipeResponse: ctx.GlobalBool("stdout"),
		Method:       "PUT",
		URL:          u.String(),
		Body:         r,
		Writer:       os.Stdout,
	}
	return Request(reqConf)
}
