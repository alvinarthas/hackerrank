package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n) // Scan input n
	var e, cp, cn, cz int

	for i := 0; i < n; i++ {
		fmt.Scan(&e)

		if e < 0 {
			cn++
		} else if e > 0 {
			cp++
		} else {
			cz++
		}
	}
	fmt.Printf("%f\n", float64(cp)/float64(n))
	fmt.Printf("%f\n", float64(cn)/float64(n))
	fmt.Printf("%f\n", float64(cz)/float64(n))
}
