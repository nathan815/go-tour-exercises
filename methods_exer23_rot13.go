package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(out []byte) (int, error) {
	n, err := r13.r.Read(out)
	if err != nil {
		return 0, err
	}
	for i, c := range out {
		var min, max byte
		if c >= 'A' && c <= 'Z' {
			min, max = 'A', 'Z'
		} else if c >= 'a' && c <= 'z' {
			min, max = 'a', 'z'
		} else {
			continue
		}
		out[i] = (c + 13) % max
		if out[i] < min {
			out[i] += min - 1
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


