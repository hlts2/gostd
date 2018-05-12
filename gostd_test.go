package gostd

import (
	"bytes"
	"os"
	"testing"
)

const MaxReaderSize = 1000

func TestNewGostd(t *testing.T) {
	got := NewGostd(os.Stdin, MaxReaderSize)

	if got == nil {
		t.Error("NewGostd(io.Reader, size) is nil")
	}
}

func TestReadLine(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "hoge\n",
			expected: "hoge",
		},
		{
			input:    "hoge\nfoo\n",
			expected: "hoge",
		},
		{
			input:    "\n",
			expected: "",
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		if gostd == nil {
			t.Error("NewGostd(io.Reader, size) gostd is error")
		}

		got := gostd.ReadLine()

		if got != test.expected {
			t.Errorf("i = %d ReadLine() expected: %s, got: %s", i, test.expected, got)
		}
	}
}
