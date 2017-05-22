package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
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

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(jsonB))
	if err != nil {
		return err
	}

	// TODO(leeola): add user supplied headers here
	//
	// defaulting to json
	req.Header.Add("Accept", "application/json")

	if err := PrintRequest(os.Stdout, req); err != nil {
		return err
	}
	fmt.Println()

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return PrintResponse(os.Stdout, res)
}
