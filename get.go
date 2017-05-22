package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/leeola/fixity/util/clijson"
	"github.com/urfave/cli"
)

func GetCmd(ctx *cli.Context) error {
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

	// this is a bad way to get a root object from
	// the clijson, but the hacked together library doesn't
	// support exporting just the map.
	//
	// TODO(leeola): add proper Object support to clijson.
	//
	// Also, http query params support multiple values for the same
	// key, so i probably want to custom write this, rather than abusing
	// the clijson library. A simple split on space would probably suffice.
	var jsonObj map[string]interface{}

	if len(jsonArgs) != 0 {
		jsonB, err := clijson.CliJson(jsonArgs)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(jsonB, &jsonObj); err != nil {
			return err
		}

		q := u.Query()
		for k, v := range jsonObj {
			q.Add(k, fmt.Sprint(v))
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest("GET", u.String(), nil)
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
