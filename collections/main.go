package main

import "fmt"

func main(){
	scores := [8]int {100,80,75,92,70,93,88,67}

	sum := 0
	var goodScores [] int

	for _, score := range scores {
		sum += score
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	average := (float64(sum) / float64(len(scores)))

	fmt.Println(average)
	fmt.Println(goodScores)



}