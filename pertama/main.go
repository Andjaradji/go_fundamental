package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("HALO Semua...!!")

	result := calculation.Add(5,7)

	fmt.Printf("my result is %d", result)
	fmt.Println()

	anotherResult := calculation.MyAdd(6,3)
	fmt.Println(anotherResult)
}
