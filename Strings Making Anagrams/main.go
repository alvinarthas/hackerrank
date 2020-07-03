package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	aString := strings.Split(a, "")
	bString := strings.Split(b, "")
	sort.Strings(aString)
	sort.Strings(bString)

	collectRestString := make(map[string]int)

	for _, aWord := range aString {
		collectRestString[aWord]++
	}

	for _, bWord := range bString {
		collectRestString[bWord]--
	}
	result := 0
	for _, v := range collectRestString {
		if v >= 0 {
			result += v
		} else {
			result -= v
		}
	}

	return int32(result)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

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
