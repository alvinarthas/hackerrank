package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t < 38 {
			fmt.Println(t)
		} else {
			if t%5 > 2 {
				// round our number
				fmt.Println(((t / 5) + 1) * 5)
			} else {
				fmt.Println(t)
			}
		}
	}
}
