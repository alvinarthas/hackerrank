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

// Complete the triplets function below.
func triplets(a []int, b []int, c []int) int {
	var tripletsCount int

	// Remove the Duplicates
	distinctA := removeDuplicates(a)
	distinctB := removeDuplicates(b)
	distinctC := removeDuplicates(c)

	// Sorts
	sort.Ints(distinctA)
	sort.Ints(distinctB)
	sort.Ints(distinctC)

	for _, v := range distinctB {
		c1 := validateIndex(distinctA, v) + 1
		c3 := validateIndex(distinctC, v) + 1

		tripletsCount += c1 * c3
	}

	return tripletsCount
}

// Remove the Duplicate
func removeDuplicates(a []int) []int {
	check := make(map[int]int)

	res := make([]int, 0)

	for _, val := range a {
		check[val] = 1
	}

	for value, _ := range check {
		res = append(res, value)
	}

	return res
}

// ValidateIndex
func validateIndex(a []int, key int) int {
	low := 0
	high := len(a) - 1
	count := -1

	for low <= high {
		mid := low + (high-low)/2

		if a[mid] <= key {
			count = mid
			low = mid + 1
		} else {
			high = mid - 1
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

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

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
