package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, o, a, count int

	arr := []int{}

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&o)
		arr = append(arr, o)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	a = arr[0]

	for i := 0; i < n; i++ {

		if a == arr[i] {
			count++
		}
	}

	fmt.Println(count)
}
