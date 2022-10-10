package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{5, 4, 9, 10, 60}
	sort.Ints(input)
	fmt.Println("angka terkecil : ", input[0])
	fmt.Println("angka terbesar : ", input[len(input)-1])

}
