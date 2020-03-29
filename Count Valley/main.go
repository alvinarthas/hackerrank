package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, seaLevel, countValley int

	fmt.Scan(&n)

	var s string
	path := make([]string, n)
	fmt.Scan(&s)

	path = strings.Split(s, "")

	for i := 0; i < n; i++ {
		if path[i] == "D" {
			seaLevel--
		} else if path[i] == "U" {
			seaLevel++
		}

		if seaLevel == 0 && path[i] == "U" {
			countValley++
		}
	}

	fmt.Println(countValley)
}
