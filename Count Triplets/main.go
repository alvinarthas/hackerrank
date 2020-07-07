package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	var currentPosition, count int64
	// Create the left and right map, for comparison
	leftHashMap := make(map[int64]int64)
	rightHashMap := make(map[int64]int64)

	// Input all the array to the map
	for _, v := range arr {
		rightHashMap[v]++
	}

	for i := 0; i < len(arr); i++ {
		currentPosition = arr[i]
		rightHashMap[currentPosition]--
		leftCount := int64(0)
		rightCount := int64(0)

		// Check if leftMap is not empty
		if _, ok := leftHashMap[currentPosition/r]; ok {
			if currentPosition%r == 0 {
				leftCount = leftHashMap[currentPosition/r]
			}
		}

		// Check if rightMap is not empty
		if _, ok := rightHashMap[currentPosition*r]; ok {
			rightCount = rightHashMap[currentPosition*r]
		}

		// left amount  * right amount
		count += leftCount * rightCount
		leftHashMap[currentPosition]++
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
