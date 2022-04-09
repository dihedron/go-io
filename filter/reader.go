package filter

import "io"

type ReaderCallback func(r *Reader, data []byte, count int)

type Reader struct {
	reader    io.Reader
	callbacks []ReaderCallback
}

// NewReader creates a new Reader, applying all the provided functional options.
func NewReader(reader io.Reader, callbacks ...ReaderCallback) *Reader {
	return &Reader{
		reader:    reader,
		callbacks: callbacks,
	}
}

// Read reads a few bytes, then calls all registered
// callback one at a time.
func (r *Reader) Read(data []byte) (int, error) {
	n, err := r.reader.Read(data)
	if err == nil {
		for _, callback := range r.callbacks {
			callback(r, data, n)
		}
	}
	return n, err
}

func (r *Reader) Close() error {
	if r, ok := r.reader.(io.ReadCloser); ok {
		return r.Close()
	}
	return nil
}
