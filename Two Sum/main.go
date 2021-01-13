package main

import (
	"fmt"
)

func main() {

	var (
		arrays = []int{3, 5, 7, 2, 15, 13}
		k      = 20
		n      = len(arrays)
	)

	fmt.Println("Find Match Sum from the array")
	fmt.Println("")
	fmt.Printf("array: %v \n", arrays)
	fmt.Printf("Sum Needed: %d \n\n", k)

	fmt.Printf("---------------- \n")

	fmt.Println("use brute force method: ")
	bruteForce(arrays, k, n)

	fmt.Printf("---------------- \n")

	fmt.Println("user pointer method: ")
	usePointer(arrays, k, n)

}

func usePointer(arrays []int, k int, n int) {
	// Complexity O(n)
	i := 0
	j := n - 1

	for i < j {

		if arrays[i]+arrays[j] == k {
			fmt.Printf("Got Match at index: %d, %d", i, j)

			return
		} else if arrays[i]+arrays[j] > k {
			j--
		} else if arrays[i]+arrays[j] < k {
			i++
		}

	}

	fmt.Println("Got No Match")

	return
}

func bruteForce(arrays []int, k int, n int) {
	// Complexity O(n2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {

			if arrays[i]+arrays[j] == k {
				fmt.Printf("Got Match at index: %d, %d\n", i, j)

				return
			}
		}
	}

	fmt.Println("Got No Match")

	return
}
