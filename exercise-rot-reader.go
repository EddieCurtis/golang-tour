package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (count int, err error) {
	count, err = reader.r.Read(b)
	if err == nil {
		length := len(b)
		for i := 0; i < length; i++ {
			value := int(b[i])
			if inRange13(value, 'A') || inRange13(value, 'a') {
				value += 13
			} else if inRange13(value, 'M') || inRange13(value, 'm') {
				value -= 13
			}
			b[i] = byte(value)
		}
	}
	return count, err
}

func inRange13(value int, start int) bool {
	return value >= start && value < start+13
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
