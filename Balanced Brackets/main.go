package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the isBalanced function below.
// The concept is, if open then push, meet the close then pop, but pop itu you find the match
func isBalanced(s string) string {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '{', '[', '(':
			stack = append(stack, char) // Push
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return "NO"
			}
			stack = stack[:len(stack)-1] // Pop
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return "NO"
			}
			stack = stack[:len(stack)-1] // Pop
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return "NO"
			}
			stack = stack[:len(stack)-1] // Pop
		}
	}
	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
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
