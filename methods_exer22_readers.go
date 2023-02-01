package main

import (
	"golang.org/x/tour/reader"
//	"fmt"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (n int, err error) {
	copy(b, []byte{'A'})
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}


