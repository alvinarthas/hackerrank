package main

import "fmt"

func fibonacciRec(n int) int {
	var result int

	if n == 1 || n == 2 {
		result = 1
	} else {
		result = fibonacci(n-1) + fibonacci(n-2)
	}

	return result
}

func main2() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacciRec(n))
}
