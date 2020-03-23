package main

import "fmt"
import "math"

func main() {
	var n int

	fmt.Scan(&n) // Scan input n

	var sum1, sum2 int
	var result float64
	var e int
	for i := 0; i < n; i++ { // make the x
		for j := 0; j < n; j++ { // make the y
			fmt.Scan(&e) // get the input per diagonal
			if i == j {
				sum1 += e
			}

			if j == n-1-i {
				sum2 += e
			}
		}
	}

	result = math.Abs(float64(sum1 - sum2)) //use math to get absolut
	fmt.Println(result)
}
