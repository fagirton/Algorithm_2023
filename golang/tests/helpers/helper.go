package helpers

import (
	"os"
	"testing"
)

func Replacer(input string, t *testing.T) (*os.File, *os.File) {
	inp := []byte(input)

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(inp)
	if err != nil {
		t.Fatal(err)
	}
	return r, w
}
