package internal_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal"
)

func TestSumm(t *testing.T) {
	assert := assert.New(t)
	input := []byte("10 12\n")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	// Restore stdin right after the test.
	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	os.Stdin = r
	os.Stdout = w

	internal.Summ()

	w.Close()
	out, _ := io.ReadAll(r)
	expected := "22\n"
	assert.Equal(expected, string(out))
}
