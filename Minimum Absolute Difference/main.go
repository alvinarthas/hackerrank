package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Abs returns the absolute value of x.
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//type foo that implements Sort for int32 type
type foo []int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Complete the minimumAbsoluteDifference function below.
func minimumAbsoluteDifference(arr []int32) int32 {
	var f foo
	f = arr
	fmt.Println(f)
	temp := Abs(f[0] - f[1])

	sort.Sort(f)
	fmt.Println(f)

	for i := 0; i < len(f)-1; i++ {
		if Abs(f[i]-f[i+1]) < temp {
			temp = Abs(f[i] - f[i+1])
		}
	}

	return temp
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

	result := minimumAbsoluteDifference(arr)

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
