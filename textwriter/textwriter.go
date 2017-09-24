package textwriter

import (
	"bytes"
	"io"
)

type TextWriter struct {
	bufSize int
	buf     bytes.Buffer
	w       io.Writer
}

func New(w io.Writer) *TextWriter {
	return &TextWriter{
		w: w,
	}
}

func (w *TextWriter) Write(p []byte) (int, error) {
	defer w.buf.Reset()

	pLen := len(p)
	if pLen > w.bufSize {
		w.buf.Grow(pLen - w.bufSize + 1)
	}

	for _, c := range p {
		if c < 32 || c > 128 {
			continue
		}

		if err := w.buf.WriteByte(c); err != nil {
			return 0, err
		}
	}

	if _, err := w.buf.WriteTo(w.w); err != nil {
		return 0, err
	}

	return pLen, nil
}
