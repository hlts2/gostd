package gostd

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unsafe"
)

// Gostd is Gostd interface
type Gostd interface {
	ReadLine() string
	ReadLineSplit(sep string) []string
	ReadLineInt() int
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

// ReadLineSplit split single-line into all substrings separated by sep
func (g *gostd) ReadLineSplit(sep string) []string {
	return strings.Split(g.ReadLine(), sep)
}

// ReadLineInt reads lines as int type
func (g *gostd) ReadLineInt() int {
	n, err := strconv.Atoi(g.ReadLine())
	if err != nil {
		panic(err)
	}

	return n
}
