package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, temp, c int

	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Sort(sort.IntSlice(arr))

	for i := 0; i < n; i++ {
		if arr[i] == temp {
			c++
			temp = 0
		} else {
			temp = arr[i]
		}
	}

	fmt.Println(c)
}
