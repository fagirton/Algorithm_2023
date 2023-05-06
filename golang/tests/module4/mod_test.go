package module4_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal/module4"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestNearestLowest(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case 1", func(t *testing.T) {
		r, w := helpers.Replacer("10\n1 2 3 2 1 4 2 5 3 1\n", t)
		os.Stdin = r
		os.Stdout = w

		module4.Nearest_Lowest()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("-1 4 3 4 -1 6 9 8 9 -1 ", string(out))
	})
}
