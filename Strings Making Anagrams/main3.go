package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// // Complete the makeAnagram function below.
// func makeAnagrama(a string, b string) int32 {
// 	chars := make(map[rune]int)

// 	for _, c := range a {
// 		chars[c]++
// 	}

// 	for _, c := range b {
// 		chars[c]--
// 	}

// 	result := 0
// 	for _, v := range chars {
// 		if v >= 0 {
// 			result += v
// 		} else {
// 			result -= v
// 		}
// 	}
// 	return int32(result)
// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 1024*1024)

// 	a := readLine(reader)

// 	b := readLine(reader)

// 	res := makeAnagrama(a, b)

// 	fmt.Fprintf(writer, "%d\n", res)

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
