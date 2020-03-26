package main

import "fmt"

func main() {
	var n, lowScore, maxScore, min, max int
	fmt.Scan(&n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i == 0 {
			lowScore = a[i]
			maxScore = a[i]
		} else {
			if a[i] < lowScore {
				lowScore = a[i]
				min++
			} else if a[i] > maxScore {
				maxScore = a[i]
				max++
			}
		}
	}

	fmt.Println(max, min)
}
