package main

import (
	// "fmt"

	"fmt"

	"isuct.ru/informatics2022/internal/module5"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// var str string
	// for scanner.Scan() {
	// 	str = str + scanner.Text()
	// }
	// fmt.Println(module4.PSPcheck(str))

	// var arr0, arr []int
	// var a int
	// fmt.Scan(&a)
	// arr0 = append(arr0, a)
	// fmt.Scan(&a)
	// arr0 = append(arr0, a)
	// for i := 0; i < arr0[0]; i++ {
	// 	fmt.Scan(&a)
	// 	arr = append(arr, a)
	// }
	// module4.MinWindowsLol(arr, arr0[1])

	tree := module5.NewTree(*module5.NewNode(3))
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	fmt.Println(tree.GetElements())
	fmt.Println(tree.Find(5))
	fmt.Println(tree.Find(0))
}
