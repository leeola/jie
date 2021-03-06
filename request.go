package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

type Config struct {
	PipeResponse bool
	Method       string
	URL          string
	Body         io.Reader
	Writer       io.Writer
}

func Request(c Config) error {
	req, err := http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return err
	}

	// TODO(leeola): add user supplied headers here
	//
	// defaulting to json
	req.Header.Add("Accept", "application/json")

	if !c.PipeResponse {
		if err := PrintRequest(c.Writer, req); err != nil {
			return err
		}
		fmt.Println()
	}

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !c.PipeResponse {
		if err := PrintResponse(c.Writer, res); err != nil {
			return err
		}
	} else {
		if err := PrintResponseBody(c.Writer, res); err != nil {
			return err
		}
	}

	return nil
}

func PrintRequest(out io.Writer, r *http.Request) error {
	urlStr := r.URL.Path
	// if the method is GET, we're going to print the querystring as the body,
	// so remove it from this url.
	if r.Method != "GET" {
	} else {
		urlStr = fmt.Sprintf("%s?%s", urlStr, r.URL.RawQuery)
	}

	c := color.New()
	c.Add(color.FgHiMagenta)
	c.Fprint(out, r.Method, " ")
	c.Add(color.FgHiCyan)
	c.Fprint(out, urlStr, " ")
	c.Add(color.FgHiMagenta)
	c.Fprintln(out, r.Proto)

	PrintHeaders(out, r.Header)

	if r.Method == "GET" && r.URL.RawQuery != "" {
		fmt.Fprintf(out, "\n?%s", r.URL.RawQuery)
	}

	if r.Body != nil {
		fmt.Println()

		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		if err := PrintJson(out, b); err != nil {
			return fmt.Errorf("invalid request json: %s", err)
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	fmt.Println()

	return nil
}
