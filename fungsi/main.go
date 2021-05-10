package main

import (
	"errors"
	"fmt"
)

func main (){
	myInputs := []int{10,7,9,8,6,3}
	total := sum(myInputs)

	fmt.Println(total)

	result, error := calculate(10,2,"*")
	
	if error == nil {
		fmt.Println("Hasil dari calculate adalah: ",result)
	} else {
		fmt.Println(error)
	}

}


func sum (inputs []int) int {
	total := 0

	for _, input := range inputs {
		total += input
	}

	return total
}

func calculate (num1,num2 int, sign string) (float64, error) {
	switch sign {
	case "+":
		return float64(num1) + float64(num2), nil
	case "-":
		return float64(num1) - float64(num2), nil
	case "*":
		return float64(num1) * float64(num2), nil
	case "/":
		return float64(num1) / float64(num2), nil
	default:
		return 0, errors.New("Sign is not acceptable")
	}
}