package module2_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module2"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestRadixSort(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case 1", func(t *testing.T) {
		r, w := helpers.Replacer("9\n12\n32\n45\n67\n98\n29\n61\n35\n09\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.RadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("[09, 12, 29, 32, 35, 45, 61, 67, 98]", string(out))
	})
	t.Run("Case: 1 1", func(t *testing.T) {
		r, w := helpers.Replacer("1 1\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.RadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("2\n", string(out))
	})
	t.Run("Case: 10000 10000", func(t *testing.T) {
		r, w := helpers.Replacer("10000 10000\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.RadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("20000\n", string(out))
	})
}

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case 1", func(t *testing.T) {
		r, w := helpers.Replacer("1\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.Merging()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("1", string(out))
	})
	t.Run("Case: 3 1", func(t *testing.T) {
		r, w := helpers.Replacer("2\n3 1\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.RadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("1 3\n", string(out))
	})
	t.Run("Case: 5 to 1", func(t *testing.T) {
		r, w := helpers.Replacer("5\n5 4 3 2 1\n", t)
		os.Stdin = r
		os.Stdout = w

		module2.RadixSort()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("1 2 3 4 5\n", string(out))
	})
}
