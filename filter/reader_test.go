package filter

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"testing"
)

const expected = "2dd6a97de8d801b4836d1aab2cadb11ff96ebd8a8b03d4f30a91debf3232401b"

func TestReader(t *testing.T) {
	file, err := os.Open("../checksum/test.json")
	if err != nil {
		t.Fatal(err)
	}

	hash := sha256.New()
	reader := NewReader(file, func(r *Reader, data []byte, count int) {
		hash.Write(data[:count])
	})
	defer reader.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, reader); err != nil {
		t.Fatal(err)
	} else {
		result := hex.EncodeToString(hash.Sum(nil))
		if result != expected {
			t.Errorf("string not equal, expected %q, got %q", expected, result)
		}
	}

}
