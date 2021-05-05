// Golang program to compare times
package main

import (
	"fmt"
	"time"
)

// importing time module

// Main function
func main() {

	var (
		testTime   time.Time
		testStrArr = []string{"hai"}
	)

	strFirstDate := "2021-05-01"
	firstDate, _ := time.Parse("2006-01-02", strFirstDate)

	strNextDate := "2021-05-03"
	nextDate, _ := time.Parse("2006-01-02", strNextDate)

	tomorrow := firstDate.Add(24 * time.Hour)

	fmt.Println(firstDate)
	fmt.Println(nextDate)
	fmt.Println(tomorrow)

	if nextDate == tomorrow {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// testArray := []int{1}

	fmt.Println(testTime)
	fmt.Println("masuk")
	for i := 0; i < len(testStrArr); i++ {
		fmt.Println(i)
	}

	// for i := 0; i < len(testArray); i++ {
	// 	fmt.Println(testArray[i+1])
	// }

	// Using time.Before() method
	// g1 := today.Before(tomorrow)
	// fmt.Println("today before tomorrow:", g1)

	// Using time.After() method
	// g2 := tomorrow.After(today)
	// fmt.Println("tomorrow after today:", g2)

}
