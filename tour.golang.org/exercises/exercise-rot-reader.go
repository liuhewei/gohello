package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	if err == nil {

		for i := 0; i < n; i++ {
			switch {
			case b[i] >= 'A' && b[i] <= 'Z':
				b[i] = 'A' + (b[i]-'A'+13)%26
			case b[i] >= 'a' && b[i] <= 'z':
				b[i] = 'a' + (b[i]-'a'+13)%26
			default:
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
