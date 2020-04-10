package main

import "fmt"

func main() {
	var q, x, y, z int

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		fmt.Scan(&z)

		catsA := x - z
		catsB := y - z

		if catsA < 0 {
			catsA *= -1
		}

		if catsB < 0 {
			catsB *= -1
		}

		if catsA > catsB {
			fmt.Println("Cat B")
		} else if catsB > catsA {
			fmt.Println("Cat A")
		} else if catsA == catsB {
			fmt.Println("Mouse C")
		}

	}
}
