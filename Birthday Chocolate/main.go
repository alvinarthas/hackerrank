package main

import "fmt"

func main() {
	var n, d, m, cake, c int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Scan(&d, &m)

	for i := 0; i+(m-1) < n; i++ {
		cake = 0
		for j := 0; j < m; j++ {
			cake += a[i+j]
		}
		if cake == d {
			c++
		}
	}

	fmt.Println(c)
}
