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
	type testCase struct {
		input    string
		isPanic  bool
		expected string
	}

	tests := []testCase{
		{
			input:    "hoge\n",
			isPanic:  false,
			expected: "hoge",
		},
		{
			input:    "hoge\nfoo\n",
			isPanic:  false,
			expected: "hoge",
		},
		{
			input:    "\n",
			isPanic:  false,
			expected: "",
		},
		{
			input:    "",
			isPanic:  true,
			expected: "",
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()
			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			if gostd == nil {
				t.Error("NewGostd(io.Reader, size) gostd is error")
			}

			got := gostd.ReadLine()

			if got != test.expected {
				t.Errorf("i = %d ReadLine() expected: %s, got: %s", i, test.expected, got)
			}

		}(i, test)
	}
}

func TestReadLineSplit(t *testing.T) {
	type testCase struct {
		input    string
		sep      string
		isPanic  bool
		expected []string
	}

	tests := []testCase{
		{
			input:    "hoge hoge\n",
			sep:      " ",
			isPanic:  false,
			expected: []string{"hoge", "hoge"},
		},
		{
			input:    "hoge,foo\n",
			sep:      ",",
			isPanic:  false,
			expected: []string{"hoge", "foo"},
		},
		{
			input:    "hoge,foo\nvar",
			sep:      ",",
			isPanic:  false,
			expected: []string{"hoge", "foo"},
		},
		{
			input:    "\n",
			sep:      "",
			isPanic:  false,
			expected: []string{},
		},
		{
			input:   "",
			sep:     "",
			isPanic: true,
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()

			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			if gostd == nil {
				t.Error("NewGostd(io.Reader, size) gostd is error")
			}

			got := gostd.ReadLineSplit(test.sep)

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("i = %d ReadLine() expected: %s, got: %s", i, test.expected, got)
			}
		}(i, test)
	}
}

func TestReadLineIntSplit(t *testing.T) {
	type testCase struct {
		input    string
		sep      string
		isPanic  bool
		expected []int
	}

	tests := []testCase{
		{
			input:    "1 2 3 4\n",
			sep:      " ",
			isPanic:  false,
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    "1\n",
			sep:      " ",
			isPanic:  false,
			expected: []int{1},
		},
		{
			input:   "1 2 3 4 \n",
			sep:     " ",
			isPanic: true,
		},
		{
			input:   "1 2 3  true\n",
			sep:     " ",
			isPanic: true,
		},
		{
			input:   "",
			sep:     " ",
			isPanic: true,
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()

			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			if gostd == nil {
				t.Error("NewGostd(io.Reader, size) gostd is error")
			}

			got := gostd.ReadLineIntSplit(test.sep)

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("i = %d ReadLineIntSplit(sep) expected: %v, got: %v", i, test.expected, got)
			}
		}(i, test)
	}
}

func TestReadLineInt(t *testing.T) {
	type testCase struct {
		input    string
		isPanic  bool
		expected int
	}

	tests := []testCase{
		{
			input:    "1\n",
			isPanic:  false,
			expected: 1,
		},
		{
			input:   "a\n",
			isPanic: true,
		},
		{
			input:   "1 \n",
			isPanic: true,
		},
		{
			input:   "1 1\n",
			isPanic: true,
		},
		{
			input:   "",
			isPanic: true,
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()

			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			if gostd == nil {
				t.Error("NewGostd(io.Reader, size) gostd is error")
			}

			got := gostd.ReadLineInt()

			if got != test.expected {
				t.Errorf("i = %d ReadLineInt() expected: %d, got: %d", i, test.expected, got)
			}
		}(i, test)
	}
}

func TestReadLineFloat64(t *testing.T) {
	type testCase struct {
		input    string
		isPanic  bool
		expected float64
	}

	tests := []testCase{
		{
			input:    "11.1\n",
			isPanic:  false,
			expected: 11.1,
		},
		{
			input:   "true\n",
			isPanic: true,
		},
		{
			input:   "11.1 \n",
			isPanic: true,
		},
		{
			input:   "",
			isPanic: true,
		},
		{
			input:   "\n",
			isPanic: true,
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()

			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			got := gostd.ReadLineFloat64()

			if got != test.expected {
				t.Errorf("i = %d ReadLineFloat64() expected: %f, got: %f", i, test.expected, got)
			}
		}(i, test)
	}
}

func TestReadLineBool(t *testing.T) {
	type testCase struct {
		input    string
		isPanic  bool
		expected bool
	}

	tests := []testCase{
		{
			input:    "true\n",
			isPanic:  false,
			expected: true,
		},
		{
			input:    "false\n",
			isPanic:  false,
			expected: false,
		},
		{
			input:   "false \n",
			isPanic: true,
		},
		{
			input:   "true \n",
			isPanic: true,
		},
		{
			input:   "11 \n",
			isPanic: true,
		},
		{
			input:   "\n",
			isPanic: true,
		},
		{
			input:   "",
			isPanic: true,
		},
	}

	for i, test := range tests {
		func(i int, test testCase) {
			defer func() {
				var isPanic bool
				if err := recover(); err != nil {
					isPanic = true
				}

				if test.isPanic != isPanic {
					t.Errorf("i = %d expected isPanic: %v, got: %v", i, test.isPanic, isPanic)
				}
			}()

			stdin := bytes.NewBufferString(test.input)

			gostd := NewGostd(stdin, MaxReaderSize)

			got := gostd.ReadLineBool()

			if got != test.expected {
				t.Errorf("i = %d ReadLineBool() expected: %v, got: %v", i, got, test.expected)
			}
		}(i, test)
	}
}
