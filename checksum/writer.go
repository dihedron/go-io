package checksum

import (
	"encoding/hex"
	"hash"
	"io"
)

type CheckSumWriter struct {
	writer   io.Writer
	checksum hash.Hash
}

func NewWriter(writer io.Writer, checksum hash.Hash) *CheckSumWriter {
	return &CheckSumWriter{
		writer:   writer,
		checksum: checksum,
	}
}

func (w *CheckSumWriter) Write(b []byte) (int, error) {
	n, err := w.writer.Write(b)
	if err == nil {
		w.checksum.Write(b[:n])
	}
	return n, err
}

func (w *CheckSumWriter) SumString() string {
	return hex.EncodeToString(w.SumBytes())
}

func (w *CheckSumWriter) SumBytes() []byte {
	return w.checksum.Sum(nil)
}

func (w *CheckSumWriter) Close() error {
	if w, ok := w.writer.(io.WriteCloser); ok {
		return w.Close()
	}
	return nil
}
