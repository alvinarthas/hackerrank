package main

import "fmt"

/*
	It turns out that golang can organize the array for queue so that the length of the array is stable
*/

func main() {
	var n, query, x int
	// Set 2 stacks for fifo(first in first out  & last in last out)
	var fifo []int
	var lifo []int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&query)
		if query == 1 {
			fmt.Scan(&x)
			fifo = append(fifo, x) // Push
			lifo = append(lifo, x) // Push
		} else if query == 2 {
			fifo = fifo[1:]
			lifo = lifo[:len(lifo)-1]
		} else if query == 3 {
			fmt.Println(fifo[0])
		}
	}
}
