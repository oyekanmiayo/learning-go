package main

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	if (len(b) == 0){
		return 0, nil
	}

	n, err = rot.r.Read(b)
	fmt.Println(b)
	
	for cnt := 0; cnt < len(b); cnt++ {
		// Uppercase
		if b[cnt] >= 'A' && b[cnt] <= 'Z' {
			// 64 + (13 - (90 - 90))
			// b[cnt] = 'A' + (13 - ('Z' - b[cnt]))
			// 76 - 65 = 11 + 13

			// b[cnt] - 65 zeros the values. i.e. we treat like values are from 0 - 26
			b[cnt]= 65 + (((b[cnt] - 65) + 13) % 26)
		}

		// Lowercase
		if b[cnt] >= 'a' && b[cnt] <= 'z' {
			b[cnt]= 97 + (((b[cnt] - 97) + 13) % 26)
		}
	}
	
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}