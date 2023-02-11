package internal

import (
	"bufio"
	"fmt"
	"os"
)

func Summ() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("You entered:", line)
}
