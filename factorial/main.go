package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(num int64) int64 {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Input : ")
	sc.Scan()
	input, _ := strconv.ParseInt(sc.Text(), 10, 64)
	fmt.Println(factorial(input))
}
