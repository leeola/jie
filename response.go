package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
)

func PrintResponse(out io.Writer, r *http.Response) error {
	defer r.Body.Close()

	c := color.New()
	c.Add(color.FgHiMagenta)
	c.Fprint(out, r.Proto, " ")
	c.Add(color.FgHiGreen)
	c.Fprintln(out, r.Status)

	PrintHeaders(out, r.Header)

	f := jsoncolor.NewFormatter()

	f.SpaceColor = color.New(color.FgRed, color.Bold)
	f.CommaColor = color.New(color.FgWhite, color.Bold)
	f.ColonColor = color.New(color.FgBlue)
	f.ObjectColor = color.New(color.FgBlue, color.Bold)
	f.ArrayColor = color.New(color.FgWhite)
	f.FieldColor = color.New(color.FgGreen)
	f.StringColor = color.New(color.FgBlack, color.Bold)
	f.TrueColor = color.New(color.FgWhite, color.Bold)
	f.FalseColor = color.New(color.FgRed)
	f.NumberColor = color.New(color.FgWhite)
	f.NullColor = color.New(color.FgWhite, color.Bold)

	resB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	prettyJson := bytes.Buffer{}
	if err := f.Format(&prettyJson, resB); err != nil {
		return err
	}

	if _, err = io.Copy(os.Stdout, &prettyJson); err != nil {
		return err
	}
	fmt.Fprintln(out)

	return nil
}
