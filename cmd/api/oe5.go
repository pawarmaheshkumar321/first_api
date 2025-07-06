package main

import "fmt"

func main() {
	N := 5
	Arr := []int{1, 2, 4, 5, 6}

	solve5(N, Arr)
}

func solve5(N int, Arr []int) int {
	totalOddNumbers := 0
	totalEvenNumbers := 0

	for _, number := range Arr {
		if number%2 == 0 {
			totalEvenNumbers++
		} else {
			totalOddNumbers++
		}
	}

	var minOps int
	if totalOddNumbers < totalEvenNumbers {
		minOps = totalOddNumbers
	} else {
		minOps = totalEvenNumbers
	}

	for _, number := range Arr {
		foundOddNumbersAtStart := 0
		foundEvenNumbersAtStart := 0

		if number%2 == 0 {
			foundEvenNumbersAtStart++
		} else {
			foundOddNumbersAtStart++
		}

		//o-e
		cost := foundOddNumbersAtStart + (totalEvenNumbers - foundEvenNumbersAtStart)
		if cost < minOps {
			minOps = cost
		}

		//e-o
		cost = foundEvenNumbersAtStart + (totalOddNumbers - foundOddNumbersAtStart)
		if cost < minOps {
			minOps = cost
		}

	}

	fmt.Printf("minops %v \n", minOps)

	return minOps
}
