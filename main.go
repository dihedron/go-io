package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dihedron/go-io/checksum"
)

func main() {

	if len(os.Args) == 1 || os.Args[1] == "--help" {
		fmt.Fprintf(os.Stderr, "%s <file> <file> <file>", os.Args[0])
	}

	for _, arg := range os.Args[1:] {
		if file, err := os.Open(arg); err != nil {
			log.Fatal(err)
		} else {
			reader := checksum.NewReader(file, sha256.New())
			defer reader.Close()

			var buffer bytes.Buffer
			if _, err := io.Copy(&buffer, reader); err == nil {
				fmt.Printf("%s  %s\n", reader.SumString(), arg)
			}
		}
	}
}
