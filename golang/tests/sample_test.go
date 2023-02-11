package internal_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"isuct.ru/informatics2022/internal"
)

func TestSumm(t *testing.T) {
	// rescueStdout := os.Stdout
	input := []byte("Alice and Bob\n")
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
	out, _ := ioutil.ReadAll(r)
	// os.Stdout = rescueStdout
	fmt.Println("Captured:", string(out))

}
