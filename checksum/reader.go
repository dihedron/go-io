package checksum

import (
	"encoding/hex"
	"hash"
	"io"
)

type CheckSumReader struct {
	reader   io.Reader
	checksum hash.Hash
}

func NewReader(reader io.Reader, checksum hash.Hash) *CheckSumReader {
	return &CheckSumReader{
		reader:   reader,
		checksum: checksum,
	}
}

func (r *CheckSumReader) Read(b []byte) (int, error) {
	n, err := r.reader.Read(b)
	if err == nil {
		r.checksum.Write(b[:n])
	}
	return n, err
}

func (r *CheckSumReader) Sum() []byte {
	return r.checksum.Sum(nil)
}

func (r *CheckSumReader) String() string {
	return hex.EncodeToString(r.checksum.Sum(nil))
}

func (r *CheckSumReader) Close() error {
	if r, ok := r.reader.(io.ReadCloser); ok {
		return r.Close()
	}
	return nil
}
