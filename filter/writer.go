package filter

import "io"

type WriterCallback func(w *Writer, data []byte, count int)

type Writer struct {
	writer    io.Writer
	callbacks []WriterCallback
}

func NewWriter(writer io.Writer, callbacks ...WriterCallback) *Writer {
	return &Writer{
		writer:    writer,
		callbacks: callbacks,
	}
}

// Write writes the given bytes, then calls all registered
// callbacks one at a time.
func (w *Writer) Write(data []byte) (int, error) {
	n, err := w.writer.Write(data)
	if err == nil {
		for _, callback := range w.callbacks {
			callback(w, data, n)
		}
	}
	return n, err
}

func (w *Writer) Close() error {
	if r, ok := w.writer.(io.WriteCloser); ok {
		return r.Close()
	}
	return nil
}
