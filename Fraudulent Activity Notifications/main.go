package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int, d int) int {
	buckets := make([]int, 201)
	notifications := 0
	for i := 0; i < len(expenditure); i++ {
		if i >= d {
			if float64(expenditure[i]) >= 2*getMedian(buckets, d) {
				notifications++
			}
		}

		buckets[expenditure[i]]++

		if i >= d {
			buckets[expenditure[i-d]]--
		}
	}

	return notifications

}

func getMedian(buckets []int, n int) float64 {
	if n%2 == 0 {
		return float64((getNth(buckets, n/2) + getNth(buckets, n/2+1))) / 2.0
	}

	return float64(getNth(buckets, n/2+1))

}
func getNth(buckets []int, n int) int {
	curr := 0

	for i := 0; i < len(buckets); i++ {
		curr += buckets[i]
		if curr >= n {
			return i
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
