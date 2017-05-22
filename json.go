package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
)

func PrintJson(out io.Writer, b []byte) error {
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

	prettyJson := bytes.Buffer{}
	if err := f.Format(&prettyJson, b); err != nil {
		return err
	}

	if _, err := io.Copy(out, &prettyJson); err != nil {
		return err
	}

	fmt.Fprintln(out)

	return nil
}
