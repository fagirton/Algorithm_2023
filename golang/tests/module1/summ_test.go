package module1_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"isuct.ru/informatics2022/internal/module1"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestSumm(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case: 10 12", func(t *testing.T) {
		r, w := helpers.Replacer("10 12\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.Summ()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("22\n", string(out))
	})
	t.Run("Case: 1 1", func(t *testing.T) {
		r, w := helpers.Replacer("1 1\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.Summ()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("2\n", string(out))
	})
	t.Run("Case: 10000 10000", func(t *testing.T) {
		r, w := helpers.Replacer("10000 10000\n", t)
		os.Stdin = r
		os.Stdout = w

		module1.Summ()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("20000\n", string(out))
	})
}
