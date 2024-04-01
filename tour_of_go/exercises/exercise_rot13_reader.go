package main

import (
	"io"
	"os"
	"strings"
)

var lookup = map[byte]byte{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'A',
	'O': 'B',
	'P': 'C',
	'Q': 'D',
	'R': 'E',
	'S': 'F',
	'T': 'G',
	'U': 'H',
	'V': 'I',
	'W': 'J',
	'X': 'K',
	'Y': 'L',
	'Z': 'M',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
}

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	if err != nil {
		return 0, err
	}

	for i := 0; i < n-1; i++ {
		found_rune, ok := lookup[b[i]]

		if ok {
			b[i] = found_rune
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
