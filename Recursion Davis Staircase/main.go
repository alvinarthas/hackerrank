package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the stepPerms function below.
func stepPerms(n int32) int32 {
	// Set the varible
	var total int32

	// set the base case, because 0 / 1 is a obvious case
	if n == 0 || n == 1 {
		return 1
	}

	// set the map, for collection the step
	hashMap := make(map[int32]int32)

	hashMap[0] = 1
	hashMap[1] = 1

	for i := int32(2); i <= n; i++ {
		total = 0

		for j := int32(1); j <= 3; j++ {
			if i-j >= 0 {
				total += hashMap[i-j]
			}
		}

		// save the current step to the map
		hashMap[i] = total
	}

	// return the latest step (n) total
	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
	}

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
