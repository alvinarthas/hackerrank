package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func query(a, b []rune) bool {
	abbr := make([][]bool, len(a)+1)

	for i := range abbr {
		abbr[i] = make([]bool, len(b)+1)
	}
	// //abbr[0][0] = true
	for i := 0; i < len(a); i++ {
		abbr[i][0] = true
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if unicode.IsLower(a[i-1]) && abbr[i-1][j] {
				abbr[i][j] = true
			}
			if abbr[i-1][j-1] && (unicode.ToUpper(a[i-1]) == b[j-1] || a[i-1] == b[j-1]) {
				abbr[i][j] = true
			}
		}
	}
	return abbr[len(a)][len(b)]
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return bufio.ScanWords(data, atEOF)
	}
	scanner.Split(split)
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	for q > 0 {
		scanner.Scan()
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		if query([]rune(a), []rune(b)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		q--
	}
}
