package main

import (
	"fmt"
)

func solve7(s1, s2 string, q int, queries [][]int) []int {
	n := len(s1)
	prefix := make([]int, n+2)
	for i := 2; i <= n; i++ {
		canSet := 0
		if i <= n {
			if s1[i-1-1] == '0' && s2[i-1-1] == '0' {
				canSet = 1
			}
		}
		if i <= n-1 {
			if s1[i-1] == '1' && s2[i+1-1] == '1' {
				canSet = 1
			}
		}
		prefix[i] = prefix[i-1] + canSet
	}
	res := make([]int, q)
	for i := 0; i < q; i++ {
		l, r := queries[i][0], queries[i][1]
		if r-l+1 < 2 {
			res[i] = 0
			continue
		}
		start := l + 1
		end := r
		if start > end {
			res[i] = 0
			continue
		}
		count := prefix[end] - prefix[start-1]
		res[i] = count
	}
	return res
}

func main() {
	s1 := "0000000"
	s2 := "0000000"
	q := 1
	queries := [][]int{{1, 7}}
	res := solve7(s1, s2, q, queries)
	for i := 0; i < q; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(res[i])
	}
	fmt.Println()
}
