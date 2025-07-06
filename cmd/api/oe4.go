package main

import "fmt"

func main() {

	N := 5
	Arr := []int{1, 2, 4, 5, 6}

	solve4(N, Arr)
	fmt.Printf("----minops %v\n", solve4(N, Arr))

}

func solve4(N int, Arr []int) int {

	total_even_numbers := 0
	total_odd_numbers := 0

	for _, number := range Arr {
		if number%2 == 0 {
			total_even_numbers++
		} else {
			total_odd_numbers++
		}
	}

	//fmt.Printf("total_odd_numbers %v, total_even_numbers %v \n", total_odd_numbers, total_even_numbers)

	var minops int
	if total_even_numbers > total_odd_numbers {
		minops = total_odd_numbers
	} else {
		minops = total_even_numbers
	}

	//fmt.Printf("minops %v \n", minops)

	found_even_numbers := 0
	found_odd_numbers := 0

	for _, number := range Arr {

		if number%2 == 0 {
			found_even_numbers++
		} else {
			found_odd_numbers++
		}

		// odd --even
		costOdd := found_odd_numbers + (total_even_numbers - found_even_numbers)

		if costOdd < minops {
			minops = costOdd
		}

		//fmt.Printf("costOdd %v, found_odd_numbers %v , total_even_numbers %v, , found_even_numbers %v\n", costOdd, found_odd_numbers, total_even_numbers, found_even_numbers)

		//even -- odd
		costEven := found_even_numbers + (total_odd_numbers - found_odd_numbers)

		//fmt.Printf("costEven %v, found_even_numbers %v , total_odd_numbers %v, , found_odd_numbers %v\n", costEven, found_even_numbers, total_odd_numbers, found_odd_numbers)

		if costEven < minops {
			minops = costEven

			//fmt.Printf("----minops %v\n", minops)
		}

	}

	return minops
}
