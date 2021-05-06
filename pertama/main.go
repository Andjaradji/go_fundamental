package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/multiplication"
)

func main() {
	fmt.Println("HALO Semua...!!")

	result := calculation.Add(5,7)

	fmt.Printf("my result is %d", result)
	fmt.Println()

	multiplyResult := multiplication.Multiply(3,4)
	fmt.Println(multiplyResult)


}
