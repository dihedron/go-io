package checksum

import (
	"bytes"
	"crypto/sha256"
	"io"
	"os"
	"testing"
)

const expected = "2dd6a97de8d801b4836d1aab2cadb11ff96ebd8a8b03d4f30a91debf3232401b"

func TestReader(t *testing.T) {

	file, err := os.Open("test.json")
	if err != nil {
		t.Fatal(err)
	}

	reader := NewReader(file, sha256.New())
	defer reader.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, reader); err != nil {
		t.Fatal(err)
	} else {
		if reader.String() != expected {
			t.Errorf("string not equal, expected %q, got %q", expected, reader.String())
		}
	}
}

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
		if writer.String() != expected {
			t.Errorf("string not equal, expected %q, got %q", expected, writer.String())
		}
	}
}
