package checksum

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {

	file, err := os.Open("test.json")
	if err != nil {
		t.Fatal(err)
	}

	var buffer bytes.Buffer
	writer := NewWriter(&buffer, sha256.New())
	defer writer.Close()

	if _, err := io.Copy(writer, file); err != nil {
		t.Fatal(err)
	} else {
		if writer.SumString() != expected {
			t.Errorf("string not equal, expected %q, got %q", expected, writer.SumString())
		}
	}
}
