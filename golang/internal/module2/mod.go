package module2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SortItem struct {
	index int
	value int
}

func Merge(a []SortItem, b []SortItem) []SortItem {
	c := []SortItem{}
	var i int = 0
	var j int = 0
	for k := 0; k < len(a)+len(b); k++ {
		if (j == len(b)) || ((i < len(a)) && (a[i].value <= b[j].value)) {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	if len(c) > 0 {
		fmt.Println(c[0].index, c[len(c)-1].index, c[0].value, c[len(c)-1].value)
	}
	return c
}

func Merge_sort(v []SortItem) []SortItem {
	if len(v) <= 1 {
		return v
	}
	left := Merge_sort(v[0 : len(v)/2])
	right := Merge_sort(v[len(v)/2:])
	return Merge(left, right)
}

func Merging() {
	var arr []SortItem
	var n, a int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, SortItem{i + 1, a})
	}
	res := Merge_sort(arr)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i].value, " ")
	}
	fmt.Println()
}

func Orders_goods() {
	var a int
	var goods_str []string
	var goods = make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	scanner.Scan()
	goods_str = strings.Split(scanner.Text(), " ")
	for i := 1; i <= len(goods_str); i++ {
		goods[i], _ = strconv.Atoi(goods_str[i-1])
	}
	scanner.Scan()

	scanner.Scan()
	goods_str = strings.Split(scanner.Text(), " ")
	for i := range goods_str {
		a, _ = strconv.Atoi(goods_str[i])
		goods[a]--
	}
	for i := 1; i <= len(goods); i++ {
		if goods[i] < 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

// func falseradixSortH(a []int, m int /*, k int*/) {
// 	// a - массив, m - кол-во разрядов, k - мощность алфавита (не нужно)
// 	for i := 0; i < m-1; i++ {
// 		var c []int //длина k-1
// 		var b []int //длина len(a)
// 		var t int
// 		for j := 0; i < len(a); i++ {
// 			c[IntIndex(a[j], i)]++
// 		}
// 		count := 0
// 		for j := 0; j < len(c); j++ {
// 			t = c[j]
// 			c[j] = count
// 			count += t
// 		}
// 		for j := 0; j < len(a); j++ {
// 			b[c[IntIndex(a[j], i)]] = a[j]
// 			c[IntIndex(a[j], i)]++
// 		}
// 		a = b
// 	}
// }

func RadixSortH(a []int) []int {
	max := math.Inf(-1)
	for i := 0; i < len(a); i++ {
		if float64(a[i]) > max {
			max = float64(a[i])
		}
	}
	for i := 1; int(max)/i > 0; i = i * 10 {
		a = countingSort(a, i)
	}
	return a
}

func countingSort(a []int, i int) []int {
	var count [10]int
	var output []int
	for j := 0; j < len(a); j++ {
		output = append(output, 0)
	}
	for j := 0; j < len(a); j++ {
		count[(a[j]/i)%10]++
	}
	for j := 1; j < 10; j++ {
		count[j] += count[j-1]
	}
	for j := len(output) - 1; j >= 0; j-- {
		output[count[(a[j]/i)%10]-1] = a[j]
		count[(a[j]/i)%10]--
	}
	return output
}

func IntIndex(a int, m int) int {
	return (a % (int(math.Pow(10, float64(m))))) / (int(math.Pow(10, float64(m))) / 10)
}

func RadixSort() {
	var arr []int
	var n, a int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	RadixSortH(arr)
	fmt.Println(arr)
}
