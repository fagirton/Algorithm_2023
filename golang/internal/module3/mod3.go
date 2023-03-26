package module3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get_rollinghash(s0 int64, s int64, h int64, p int64, x int64, l int64) int64 {
	return (h*x - s0*Pow(x, l) + s) % p
}

func Pow(x int64, y int64) int64 {
	for i := 0; i < int(y); i++ {
		x = x * x
	}
	return x
}

func Get_hash(s string, l int, p int64, x int64) int64 {
	res := int64(0)
	for i := 0; i < l; i++ {
		res += int64(s[i]) * Pow(x, int64(l-i))
	}
	return res % p
}

func find_with_hash(search_obj string, search_str string) {
	p := int64(1e9 + 7)
	x := int64(26) //677
	l := len(search_obj)
	search_obj_hash := Get_hash(search_obj, l, p, x)
	hash := Get_hash(search_str, l, p, x)
	for i := 1; i < len(search_str)-l; i++ {
		if hash == search_obj_hash {
			fmt.Println(i - 1)
		}
		hash = Get_rollinghash(int64(search_str[i-1]), int64(search_str[i+l]), hash, p, x, int64(l))
	}
}

func HashSearch() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	search_str := scanner.Text()
	scanner.Scan()
	search_obj := scanner.Text()
	find_with_hash(search_obj, search_str)
}

// p := int64(1e9 + 7)
// 	x := int64(26) //677
// 	res := int64(0)
// 	for i := 0; i < len(s); i++ {
// 		res = (res * x * int64(s[i])) % p
// 	}
// 	return res

func prefix_calc(s string) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s)-1; i++ {
		j := res[i]
		for j > 0 && s[i+1] != s[j] {
			j = res[j-1]
		}
		if s[i+1] == s[j] {
			res[i+1] = j + 1
		} else {
			res[i+1] = 0
		}
	}
	return res
}

func prefix_search(t string, s string) []int {
	ts := t + "&" + s
	p := prefix_calc(ts)
	var res []int
	for i := 2 * len(t); i < len(ts); i++ {
		if p[i] == len(t) {
			res = append(res, i-2*len(t))
		}
	}
	return res
}

func AlgKunt_Morris_Pratt() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	search_str := scanner.Text()
	scanner.Scan()
	search_obj := scanner.Text()
	fmt.Println(strings.Trim(fmt.Sprint(prefix_search(search_obj, search_str)), "[]"))
}

func Shift_1(s string) string {
	return string(s[len(s)-1]) + s[0:len(s)-1]
}

func loop_shift(orig string, shifted string) int {
	for i := 0; i < len(shifted); i++ {
		if orig == shifted {
			return i
		}
		shifted = Shift_1(shifted)
	}
	return -1
}

func MainLoopShift() {
	scanner := bufio.NewReader(os.Stdin)

	shifted, _ := scanner.ReadString('\n')
	original, _ := scanner.ReadString('\n')
	fmt.Println(loop_shift(original, shifted))
}

func Z_func(s string) []int {
	z := make([]int, len(s))
	l := 0
	r := 0
	for i := 1; i < len(s); i++ {
		if i <= r {
			if r-i+1 <= z[i-l] {
				z[i] = r - i + 1
			} else {
				z[i] = z[i-l]
			}
		}
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func FindMax(s []int) int {
	max := -1
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
