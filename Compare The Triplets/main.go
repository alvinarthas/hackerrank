package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int

	var alice, bob int

	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}

	fmt.Println(alice, bob)
}
