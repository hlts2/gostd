package gostd

import (
	"bytes"
	"os"
	"reflect"
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

func TestReadLineSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{
			input:    "hoge hoge\n",
			sep:      " ",
			expected: []string{"hoge", "hoge"},
		},
		{
			input:    "hoge,foo\n",
			sep:      ",",
			expected: []string{"hoge", "foo"},
		},
		{
			input:    "hoge,foo\nvar",
			sep:      ",",
			expected: []string{"hoge", "foo"},
		},
		{
			input:    "\n",
			sep:      "",
			expected: []string{},
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		if gostd == nil {
			t.Error("NewGostd(io.Reader, size) gostd is error")
		}

		got := gostd.ReadLineSplit(test.sep)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("i = %d ReadLine() expected: %s, got: %s", i, test.expected, got)
		}
	}
}

func TestReadLineIntSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []int
	}{
		{
			input:    "1 2 3 4\n",
			sep:      " ",
			expected: []int{1, 2, 3, 4},
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		if gostd == nil {
			t.Error("NewGostd(io.Reader, size) gostd is error")
		}

		got := gostd.ReadLineIntSplit(test.sep)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("i = %d ReadLineIntSplit(sep) expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestReadLineInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "1\n",
			expected: 1,
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		if gostd == nil {
			t.Error("NewGostd(io.Reader, size) gostd is error")
		}

		got := gostd.ReadLineInt()

		if got != test.expected {
			t.Errorf("i = %d ReadLineInt() expected: %d, got: %d", i, test.expected, got)
		}
	}
}

func TestReadLineFloat64(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{
			input:    "11.1\n",
			expected: 11.1,
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		got := gostd.ReadLineFloat64()

		if got != test.expected {
			t.Errorf("i = %d ReadLineFloat64() expected: %f, got: %f", i, test.expected, got)
		}
	}
}

func TestReadLineBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "true\n",
			expected: true,
		},
		{
			input:    "false\n",
			expected: false,
		},
	}

	for i, test := range tests {
		stdin := bytes.NewBufferString(test.input)

		gostd := NewGostd(stdin, MaxReaderSize)

		got := gostd.ReadLineBool()

		if got != test.expected {
			t.Errorf("i = %d ReadLineBool() expected: %v, got: %v", i, got, test.expected)
		}
	}
}
