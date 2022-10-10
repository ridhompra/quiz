package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Your name : ")
	sc.Scan()
	input := sc.Text()
	if input == "" {
		fmt.Println("NAMA HARUS DI ISI !!")
	}
	fmt.Print("Pilih peran mu \n\t1. Penyihir\n\t2. Guard \n\t3. Warewolf\n Masukan angka : ")
	sc.Scan()
	input2 := sc.Text()
	switch input2 {
	case "1":
		input2 = "Penyihir"
	case "2":
		input2 = "Guard"
	case "3":
		input2 = "Warewolf"
	default:
		fmt.Println("Pilih peranmu untuk memulai game")
		return
	}
	fmt.Println("\n====================================")
	fmt.Printf("Your name is %q\n", input)
	fmt.Printf("You choice %q\n", input2)

}
