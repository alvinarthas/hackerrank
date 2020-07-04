package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	// Devide the luck and important
	var important []int
	var luck int
	for i := 0; i < n; i++ {
		var l, t int
		fmt.Scanf("%d %d", &l, &t)
		// not important, so just add the luck
		if t == 0 {
			luck += l
		} else { // oh it is important, so let's hold it
			important = append(important, l)
		}
	}
	sort.Ints(important)
	for i := len(important) - 1; i >= 0; i-- {
		if k > 0 {
			luck += important[i]
			k--
		} else {
			luck -= important[i]
		}
	}
	fmt.Println(luck)
}
