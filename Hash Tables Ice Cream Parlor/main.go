package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
	hashMap := make(map[int32][]int)
	for s, v := range cost {
		moneyNeed := money - v

		// Look, if the difference cost available in the hashMap
		if j, ok := hashMap[moneyNeed]; ok {
			// Check if the index whether still in the begining or not
			if s <= j[0] {
				fmt.Printf("%d %d\n", s+1, j[0]+1)
			} else {
				fmt.Printf("%d %d\n", j[0]+1, s+1)
			}

			// return immediately, because we already find the pair
			return
		}

		// if not available, insert to hashMap
		hashMap[v] = append(hashMap[v], s)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
