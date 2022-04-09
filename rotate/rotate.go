package rotate

import (
	"io"
	"sync"
)

type Writer struct {
	lock      sync.RWMutex
	writer    io.Writer
	chunksize int
	total     int
}

func NewWriter(writer io.Writer, chunksize int) *Writer {
	return &Writer{
		writer:    writer,
		chunksize: chunksize,
	}
}

func (w *Writer) Write(data []byte) (int, error) {
	/*
		s := string(data)
		if strings.HasSuffix(s, "\n") {

		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {

		}
	*/
	n, err := w.writer.Write(data)
	return n, err
}

func (w *Writer) Close() error {
	if r, ok := w.writer.(io.WriteCloser); ok {
		return r.Close()
	}
	return nil
}
