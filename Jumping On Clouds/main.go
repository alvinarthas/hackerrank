package main

import "fmt"

func main() {
	var n, count, j int
	fmt.Scan(&n)

	c := make([]int, n)

	// scan the slice
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}

	// Loop until break
	for {
		if j+2 < len(c) && c[j+2] == 0 {
			j += 2
		} else if j+1 < len(c) {
			j++
		} else {
			break
		}
		count++
	}

	fmt.Println(count)
}
