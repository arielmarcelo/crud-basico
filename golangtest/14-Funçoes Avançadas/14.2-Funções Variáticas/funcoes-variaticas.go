package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(12, 15, 15, 17, 165, 15, 58, 13, 18)
	fmt.Println(totalDaSoma)
}
