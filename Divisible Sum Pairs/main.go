package main

import "fmt"

func main() {
	var n, k, c int

	fmt.Scan(&n, &k)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (arr[i]+arr[j])%k == 0 {
				c++
			}
		}
	}

	fmt.Println(c)
}
