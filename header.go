package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func PrintHeaders(out io.Writer, h http.Header) {
	for k, v := range h {
		c := color.New()
		c.Add(color.FgHiBlack)
		c.Fprint(out, k)
		c.Add(color.FgWhite)
		c.Fprint(out, ": ")
		c.Add(color.FgCyan)
		c.Fprintf(out, "%s\n", strings.Join(v, ", "))
	}
}
