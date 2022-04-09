package checksum

import (
	"encoding/hex"
	"hash"
	"io"

	"github.com/dihedron/go-io/filter"
)

type Writer struct {
	filter.Writer
	checksum hash.Hash
}

func NewWriter(writer io.Writer, checksum hash.Hash) *Writer {
	return &Writer{
		Writer: *filter.NewWriter(writer, func(w *filter.Writer, data []byte, count int) {
			checksum.Write(data[:count])
		}),
		checksum: checksum,
	}
}

func (w *Writer) SumString() string {
	return hex.EncodeToString(w.SumBytes())
}

func (w *Writer) SumBytes() []byte {
	return w.checksum.Sum(nil)
}
