package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the isValid function below.
func isValid(s string) string {
	var sameAsMin, sameAsMax int

	hashMap := make(map[rune]int)

	// Collect to HashMap
	for _, c := range s {
		hashMap[c]++
	}

	// Cannot proceed if size more than 2
	if checkSize(hashMap) == 1 {
		return "YES"
	} else if checkSize(hashMap) > 2 {
		return "NO"
	}

	// Get the max and min value
	max, min := findMinMax(hashMap)

	// count how many sum same with min and max
	for _, v := range hashMap {
		if v == min {
			sameAsMin++
		} else if v == max {
			sameAsMax++
		}
	}

	if min == 1 && sameAsMin == 1 {
		return "YES"
	} else if max-min == 1 && sameAsMax == 1 {
		return "YES"
	}

	return "NO"
}

// Find Max Value from the Map
func findMinMax(hashMap map[rune]int) (int, int) {
	var max, min int

	for _, v := range hashMap {
		if v > max {
			max = v
		} else if v < max {
			min = v
		}
	}

	return max, min
}

// CheckSize
func checkSize(hashMap map[rune]int) int {
	tempMap := make(map[int]int)
	var count int

	for _, v := range hashMap {
		if tempMap[v] < 1 {
			tempMap[v]++
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
