package module1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Summ() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	arr := strings.Split(line, " ")

	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])
	summ := a + b
	fmt.Println(summ)
}
