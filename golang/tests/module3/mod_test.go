package module3_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal/module3"
	"isuct.ru/informatics2022/tests/helpers"
)

func TestHash(t *testing.T) {
	assert := assert.New(t)

	defer func(v *os.File) { os.Stdin = v }(os.Stdin)
	defer func(v *os.File) { os.Stdout = v }(os.Stdout)
	t.Run("Case 1", func(t *testing.T) {
		r, w := helpers.Replacer("ababbababa\naba\n", t)
		os.Stdin = r
		os.Stdout = w

		module3.HashSearch()
		w.Close()
		out, _ := io.ReadAll(r)
		assert.Equal("0 5 7", string(out))
	})

	t.Run("Test Rolling hash", func(t *testing.T) {
		var search_obj string = "silly"
		var search_new string = "illyx"
		p := int64(1e9 + 7)
		x := int64(26) //677
		l := len(search_obj)
		hash := module3.Get_hash(search_obj, l, p, x)
		s := "sx"
		assert.Equal(module3.Get_hash(search_new, l, p, x), module3.Get_rollinghash(int64(s[0]), int64(s[1]), hash, p, x, int64(l)))
	})
}
