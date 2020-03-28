package main

import "fmt"

func main() {
	var n, p, turnLtoR, totalLtoR, turnRtoL int

	fmt.Scan(&n, &p)

	totalLtoR = n / 2

	turnLtoR = p / 2

	turnRtoL = totalLtoR - turnLtoR

	if turnRtoL < turnLtoR {
		fmt.Println(turnRtoL)
	} else {
		fmt.Println(turnLtoR)
	}
}
