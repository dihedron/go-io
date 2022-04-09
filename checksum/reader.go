package checksum

import (
	"encoding/hex"
	"hash"
	"io"

	"github.com/dihedron/go-io/filter"
)

type Reader struct {
	filter.Reader
	checksum hash.Hash
}

func NewReader(reader io.Reader, checksum hash.Hash) *Reader {
	return &Reader{
		Reader: *filter.NewReader(reader, func(w *filter.Reader, data []byte, count int) {
			checksum.Write(data[:count])
		}),
		checksum: checksum,
	}
}

func (r *Reader) SumString() string {
	return hex.EncodeToString(r.SumBytes())
}

func (r *Reader) SumBytes() []byte {
	return r.checksum.Sum(nil)
}
