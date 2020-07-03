package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	bottomUp := make([]int, n+1)
	bottomUp[1] = 1
	bottomUp[2] = 1

	for i := 3; i <= n; i++ {
		bottomUp[i] = bottomUp[i-1] + bottomUp[i-2]
	}
	return bottomUp[n]
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
