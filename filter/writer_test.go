package filter

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {
	file, err := os.Open("../checksum/test.json")
	if err != nil {
		t.Fatal(err)
	}

	hash := sha256.New()
	var buffer bytes.Buffer
	writer := NewWriter(&buffer, func(r *Writer, data []byte, count int) {
		hash.Write(data[:count])
	})
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		t.Fatal(err)
	} else {
		result := hex.EncodeToString(hash.Sum(nil))
		if result != expected {
			t.Errorf("string not equal, expected %q, got %q", expected, result)
		}
	}
}
