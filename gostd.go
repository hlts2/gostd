package gostd

import (
	"bufio"
	"io"
	"unsafe"
)

// Gostd is Gostd interface
type Gostd interface {
	ReadLine() string
}

type gostd struct {
	readerSize int
	reader     *bufio.Reader
}

// NewGostd returns gostd instance
func NewGostd(reader io.Reader, readerSize int) Gostd {
	return &gostd{
		readerSize: readerSize,
		reader:     bufio.NewReaderSize(reader, readerSize),
	}
}

// ReadLine returns single-line
// The text returned from ReadLine does not include the line end ("\r\n" or "\n")
func (g *gostd) ReadLine() string {
	buf := make([]byte, 0, g.readerSize)

	for {
		line, isPrefix, err := g.reader.ReadLine()
		if err != nil {
			panic(err)
		}

		buf = append(buf, line...)
		if !isPrefix {
			break
		}
	}

	return *(*string)(unsafe.Pointer(&buf))
}
