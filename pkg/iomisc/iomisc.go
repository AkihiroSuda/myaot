package iomisc

import (
	"io"
)

// ForceWriter panics on an error.
func ForceWriter(w io.Writer) io.Writer {
	return &forceWriter{
		Writer: w,
	}
}

type forceWriter struct {
	io.Writer
}

func (w *forceWriter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	if err != nil {
		panic(err)
	}
	return n, nil
}
