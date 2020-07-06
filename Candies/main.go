package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the candies function below.
func candies(n int32, arr []int32, candyLeft []int32, candyRight []int32) int64 {
	var count int64

	// Set total candy based from left or right
	for i := int32(1); i < n; i++ {
		// for counting from right to left
		j := n - (i + 1)

		if arr[i] > arr[i-1] {
			candyLeft[i] = candyLeft[i-1] + 1
		}

		if arr[j] > arr[j+1] {
			candyRight[j] = candyRight[j+1] + 1
		}
	}

	// Compare each array, get the max
	for i := int32(0); i < n; i++ {
		count += max(candyLeft[i], candyRight[i])
	}

	return count
}

// Get the max
func max(candyLeft int32, candyRight int32) int64 {
	if candyLeft > candyRight {
		return int64(candyLeft)
	}
	return int64(candyRight)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr, candyLeft, candyRight []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
		candyLeft = append(candyLeft, 1)
		candyRight = append(candyRight, 1)
	}

	result := candies(n, arr, candyLeft, candyRight)

	fmt.Fprintf(writer, "%d\n", result)

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
