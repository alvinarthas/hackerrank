package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	// Initialize variable
	isVisited := make([]bool, len(arr)+1) //to check visited variable, start from false
	nodeMap := make(map[int32]int32)
	var count, temp int32

	// Insert the arr data to the hashmap start from index 1
	for i := 1; i <= len(arr); i++ {
		nodeMap[int32(i)] = arr[i-1]
	}

	// Loop over start from index 1
	for j := 1; j <= len(arr); j++ {

		// Check if the index isn't visited
		if isVisited[j] == false {
			isVisited[j] = true                // visit the index
			if int32(j) != nodeMap[int32(j)] { // check if the value is not equal
				temp = nodeMap[int32(j)]

				for !isVisited[temp] {
					isVisited[temp] = true
					temp = nodeMap[temp]
					count++ // place count for the swap
				}
			}
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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
