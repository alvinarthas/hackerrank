package main

import "fmt"

func main() {
	var n, b, m int
	current := -1

	fmt.Scan(&b, &n, &m)

	keyboards := make([]int, n)
	drives := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&keyboards[i])
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&drives[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum := keyboards[i] + drives[j]
			if sum <= b && sum > current {
				current = sum
			}
		}
	}

	fmt.Println(current)
}
