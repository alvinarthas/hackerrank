package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 5)
	var sumMin, sumMax int
	fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4])

	sort.Ints(a)

	for i := 0; i < 4; i++ {
		sumMin += a[i]
	}
	for i := 1; i < 5; i++ {
		sumMax += a[i]
	}

	fmt.Println(sumMin, sumMax)
}
