package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b *byte) {
	switch true {
	case (*b >= 'A' && *b <= 'M') || (*b >= 'a' && *b <= 'm'):
		*b += 13
	case (*b >= 'N' && *b <= 'Z') || (*b >= 'n' && *b <= 'z'):
		*b -= 13
	}
}

func (rot13Reader *rot13Reader) Read(p []byte) (int, error) {
	n, err := rot13Reader.r.Read(p)

	for i := range p {
		rot13(&p[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
