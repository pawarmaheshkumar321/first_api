package main

import "fmt"

func solve2(N int, Arr []int) int {
	// Precompute prefix and suffix arrays for odd and even costs
	prefixOdd := make([]int, N+1) // prefixOdd[i] is cost to make first i elements odd
	for i := 1; i <= N; i++ {
		cost := 0
		if Arr[i-1]%2 == 0 {
			cost = 1
		}
		prefixOdd[i] = prefixOdd[i-1] + cost
	}

	fmt.Println("prefixOdd :", prefixOdd)

	suffixOdd := make([]int, N+2) // suffixOdd[i] is cost to make elements from i to N odd
	for i := N; i >= 1; i-- {
		cost := 0
		if Arr[i-1]%2 == 0 {
			cost = 1
		}
		suffixOdd[i] = suffixOdd[i+1] + cost
	}

	fmt.Println("suffixOdd :", suffixOdd)

	prefixEven := make([]int, N+1) // prefixEven[i] is cost to make first i elements even
	for i := 1; i <= N; i++ {
		cost := 0
		if Arr[i-1]%2 != 0 {
			cost = 1
		}
		prefixEven[i] = prefixEven[i-1] + cost
	}

	fmt.Println("prefixEven :", prefixEven)

	suffixEven := make([]int, N+2) // suffixEven[i] is cost to make elements from i to N even
	for i := N; i >= 1; i-- {
		cost := 0
		if Arr[i-1]%2 != 0 {
			cost = 1
		}
		suffixEven[i] = suffixEven[i+1] + cost
	}

	fmt.Println("suffixEven :", suffixEven)

	minCost := 1 << 30

	fmt.Println("minCost :", minCost)

	// Case 1: all odd
	case1 := prefixOdd[N]
	if case1 < minCost {
		minCost = case1
	}

	// Case 2: all even
	case2 := prefixEven[N]
	if case2 < minCost {
		minCost = case2
	}

	// Case 3: some i to j is even, rest are odd
	// The even block can be any i to j, including i=1 to j=N (case2)
	// So we need to find i and j such that:
	// cost = (cost to make 1..i-1 odd) + (cost to make i..j even) + (cost to make j+1..N odd)
	// To compute this efficiently, note that:
	// cost(i,j) = prefixOdd[i-1] + (prefixEven[j] - prefixEven[i-1]) + suffixOdd[j+1]
	// We can iterate over possible i and j, but O(N^2) is not feasible for N=1e6.
	// Wait, but the even block can be any contiguous segment. So for each possible i, the optimal j is i to N or similar.
	// Alternatively, the problem can be viewed as finding the minimal (prefixOdd[i] + suffixOdd[j+1] + (prefixEven[j] - prefixEven[i]))
	// But this seems tricky to compute in O(N).
	// Alternatively, the minimal cost for case3 is the minimal of:
	// (prefixOdd[i] + suffixEven[i+1]) for any i (split before i+1: first i elements are odd, rest are even)
	// (prefixEven[j] + suffixOdd[j+1]) for any j (first j elements even, rest odd)
	// So we can compute these in O(N) time.

	// Check all possible split points for case3
	for i := 0; i <= N; i++ {
		// First i elements are odd, rest are even
		cost := prefixOdd[i] + suffixEven[i+1]
		if cost < minCost {
			minCost = cost
		}
		// First i elements are even, rest are odd
		cost = prefixEven[i] + suffixOdd[i+1]
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func main() {
	// Sample input
	N := 5
	Arr := []int{1, 2, 4, 5, 6}
	fmt.Println(solve(N, Arr))
}
