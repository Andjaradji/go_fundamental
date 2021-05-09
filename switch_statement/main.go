package main

import "fmt"

func main() {
	number := 8

	// switch  number {
	// case 1,3,5:
	// 	fmt.Println("Angka adalah bilangan ganjil diantara satu dan 5")
	// case 8:
	// 	fmt.Println("Angka adalah 8")
	// default:
	// 	fmt.Println("Ini adalah default case")
	// }

	// switch with if else style

	switch {
	case (number < 5) && (number % 2 == 1):
		fmt.Println("Angka lebih kecil dari 5 dan ganjil")
	case  (number > 6) && (number %2 ==0 ):
		fmt.Println("Angka lebih besar dari 6 dan genap")
	default:
		fmt.Println("Ini adalah default case")
	}
}
