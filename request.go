package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

func PrintRequest(out io.Writer, r *http.Request) error {
	urlStr := r.URL.Path
	// if the method is GET, we're going to print the querystring as the body,
	// so remove it from this url.
	if r.Method != "GET" {
	} else {
		urlStr = urlStr + r.URL.RawQuery
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
			return err
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	fmt.Println()

	return nil
}
