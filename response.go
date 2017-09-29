package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

func PrintResponse(out io.Writer, r *http.Response) error {
	defer r.Body.Close()

	c := color.New()
	c.Add(color.FgHiMagenta)
	c.Fprint(out, r.Proto, " ")
	c.Add(color.FgHiGreen)
	c.Fprintln(out, r.Status)

	PrintHeaders(out, r.Header)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := PrintJson(out, b); err != nil {
		return err
	}

	return nil
}

func PrintResponseBody(out io.Writer, r *http.Response) error {
	if r.Body == nil {
		return nil
	}

	defer r.Body.Close()
	bodyB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var outB bytes.Buffer
	if err := json.Indent(&outB, bodyB, "", "  "); err != nil {
		return fmt.Errorf("invalid response json: %s", err)
	}

	if _, err := io.Copy(out, &outB); err != nil {
		return err
	}

	return nil
}
