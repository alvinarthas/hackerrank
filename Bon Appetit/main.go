package main

import "fmt"

func main() {
	var n, k, bAct, bChrg int

	fmt.Scan(&n, &k)

	bill := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&bill[i])
		if i != k {
			bChrg += bill[i]
		}
	}

	bChrg = bChrg / 2

	fmt.Scan(&bAct)

	if bAct == bChrg {
		fmt.Println("Bon Appetit")
	} else if bAct > bChrg {
		fmt.Println(bAct - bChrg)
	}
}
