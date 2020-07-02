package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	// Initial Variable
	qtyStart := int64(0)
	resultCount := int64(0)

	lenStart := len(s)
	ss := strings.Split(s, "")
	position := -1

	// Early count for a's
	for i := 0; i < lenStart; i++ {
		if ss[i] == "a" {
			if position == -1 {
				position = i
			}
			qtyStart++
		}
	}

	// if there is don't have any a's
	if position == -1 {
		return resultCount
	}

	// calculate the a's based on the qtyStart
	modulo := n % int64(lenStart)
	total := n / int64(lenStart)
	if modulo == 0 {
		resultCount = total * qtyStart
	} else {
		resultCount = total * qtyStart
		for i := int64(0); i < modulo; i++ {
			if ss[i] == "a" {
				resultCount++
			}
		}
	}
	return int64(resultCount)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
