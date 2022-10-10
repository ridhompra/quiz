package main

import "fmt"

func main() {
	fizzbuzz(15)
}
func fizzbuzz(input int) {
	for i := 1; i <= input; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
