// Golang program to compare times
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// importing time module

// Main function
func mainxx() {

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

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func main() {
	amap := make(map[int]map[int]map[int]string)

	amap[1000] = make(map[int]map[int]string)
	amap[1000][1] = map[int]string{}
	amap[1000][1][90000] = "2021-06-01,2021-06-02"

	for k1, v1 := range amap {
		for k2, v2 := range v1 {
			for k3, v3 := range v2 {
				fmt.Println(k1, k2, k3, v3)
			}
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
