package main

import "fmt"

func main(){
	// point := 8

	// if point <= 5 {
	// 	fmt.Println("Maaf Anda Tidak Lulus")
	// } else if point <= 6 {
	// 	fmt.Println("Anda lulus pas-pasan")
	// } else if point <= 7 {
	// 	fmt.Println("Anda Lulus Cukup Baik")
	// } else if point <= 8 {
	// 	fmt.Println("Anda Lulus dengan Sangat Baik")
	// } else {
	// 	fmt.Println("Anda Luar Biasa")
	// }

	// using point as temporary variable

	if point :=3; point <=5 {
		fmt.Println("Maaf Anda Tidak Lulus")
	} else if point <= 6 {
		fmt.Println("Anda lulus pas-pasan")
	} else if point <= 7 {
		fmt.Println("Anda Lulus Cukup Baik")
	} else if point <= 8 {
		fmt.Println("Anda Lulus dengan Sangat Baik")
	} else {
		fmt.Println("Anda Luar Biasa")
	}


}