package main

import "fmt"

func main() {
	// Sample input
	N := 5
	Arr := []int{1, 2, 4, 5, 6}
	fmt.Println(solve3(N, Arr))
}

func solve3(N int, Arr []int) int {

	total_odd_nums := 0
	total_even_nums := 0

	for _, number := range Arr {
		if number%2 == 0 {
			total_even_nums++
		} else {
			total_odd_nums++
		}
	}

	fmt.Println("totalOdd :", total_odd_nums)
	fmt.Println("totalEven :", total_even_nums)

	var minops int
	if total_odd_nums > total_even_nums {
		minops = total_even_nums
	} else {
		minops = total_odd_nums
	}

	prefixEven := 0
	prefixOdd := 0
	for i := 0; i < N; i++ {
		if Arr[i]%2 == 0 {
			prefixEven++
		} else {
			prefixOdd++
		}
		// Cost to make first i+1 elements odd and rest even  	---odd--evn
		cost := (prefixOdd + (total_even_nums - prefixEven))
		fmt.Printf("number : %v ,Cost: %v,  prefixOdd : %v, totalEven : %v, prefixEven : %v \n", Arr[i], cost, prefixOdd, total_even_nums, prefixEven)

		if cost < minops {
			minops = cost
		}
		// Cost to make first i+1 elements even and rest odd 	--even--odd
		cost = (prefixEven + (total_odd_nums - prefixOdd))
		fmt.Printf("number : %v , Cost2: %v, prefixEven : %v, total_odd_nums : %v, prefixOdd : %v \n", Arr[i], cost, prefixEven, total_odd_nums, prefixOdd)

		fmt.Printf("cost : %v, minOperations : %v\n", cost, minops)

		if cost < minops {
			minops = cost
		}
	}

	return minops

}
