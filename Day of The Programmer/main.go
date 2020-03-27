package main

import "fmt"

func main() {
	var y int

	fmt.Scan(&y)
	// Check if it is Leap Year
	if y < 1918 {
		if y%4 == 0 {
			fmt.Print("12.09.", y)
		} else {
			fmt.Print("13.09.", y)
		}
	} else if y > 1918 {
		if y%400 == 0 || y%4 == 0 && y%100 != 0 {
			fmt.Print("12.09.", y)
		} else {
			fmt.Print("13.09.", y)
		}
	} else {
		fmt.Print("26.09.", y)
	}
}
