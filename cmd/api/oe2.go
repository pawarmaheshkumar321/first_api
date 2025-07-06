package main

import "fmt"

func main() {
	// Sample input
	N := 5
	Arr := []int{1, 2, 4, 5, 6}
	fmt.Println(solve(N, Arr))
}
func solve(N int, Arr []int) int {
	// Count total odd and even numbers
	totalOdd := 0
	totalEven := 0
	for _, num := range Arr {
		if num%2 == 0 {
			totalEven++
		} else {
			totalOdd++
		}
	}

	fmt.Println("totalOdd :", totalOdd)
	fmt.Println("totalEven :", totalEven)

	// Case 1: Make all numbers odd (convert evens)
	minOperations := totalEven

	fmt.Println("minOperations :", minOperations)

	// Case 2: Make all numbers even (convert odds)
	if totalOdd < minOperations {
		minOperations = totalOdd
	}

	fmt.Println("minOperations2 :", minOperations)

	// Case 3: Create single even block
	// Track the best split point for even block
	prefixEven := 0
	prefixOdd := 0
	for i := 0; i < N; i++ {
		if Arr[i]%2 == 0 {
			prefixEven++
		} else {
			prefixOdd++
		}
		// Cost to make first i+1 elements odd and rest even
		cost := (prefixOdd + (totalEven - prefixEven))
		fmt.Printf("Cost: %v, i : %v , prefixOdd : %v, totalEven : %v, prefixEven : %v\n", cost, prefixOdd, totalEven, prefixEven)

		if cost < minOperations {
			minOperations = cost
		}
		// Cost to make first i+1 elements even and rest odd
		cost = (prefixEven + (totalOdd - prefixOdd))
		fmt.Printf("Cost2: %v, i : %v , prefixEven : %v, totalOdd : %v, prefixOdd : %v", cost, prefixEven, totalOdd, prefixOdd)

		fmt.Printf("cost : %v, minOperations : %v\n", cost, minOperations)

		if cost < minOperations {
			minOperations = cost
		}
	}

	return minOperations
}
