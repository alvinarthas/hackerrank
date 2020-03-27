package main

import "fmt"

func main() {
	var n, temp int

	fmt.Scan(&n)

	a := make([]int, 5)

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		temp = temp - 1
		a[temp]++
	}

	max, curInt := a[0], 0

	for i, num := range a { // ini adalah foreach
		if num > max {
			max = num
			curInt = i + 1
		}
	}
	fmt.Println(curInt)
}
