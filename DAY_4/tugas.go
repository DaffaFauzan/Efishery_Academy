package main

import (
	"fmt"
)

func star(number int) {
	// Percabangan untuk bilangan genap
	if number%2 == 0 {
		for i := 0; i < number; i++ {
			for j := 0; j < number; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		//Percabangan untuk bilangan ganjil
	} else {
		for i := 0; i < number+1; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}
func main() {
	var number int
	fmt.Scanf("%d", &number)
	star(number)
}
