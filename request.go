package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

func PrintRequest(out io.Writer, r *http.Request) error {
	var urlStr string
	// if the method is GET, we're going to print the querystring as the body,
	// so remove it from this url.
	if r.Method == "GET" {
		u := *r.URL
		u.RawQuery = ""
		urlStr = u.String()
	} else {
		urlStr = r.URL.String()
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

	fmt.Println()

	return nil
}
