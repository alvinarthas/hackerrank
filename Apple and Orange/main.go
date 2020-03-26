package main

import "fmt"

func main() {
	var s, t, a, b, m, n, mval, nval, mcount, ncount int

	fmt.Scan(&s, &t, &a, &b, &m, &n)
	for i := 0; i < m; i++ {
		fmt.Scan(&mval)

		if (a+mval) >= s && (a+mval) <= t {
			mcount++
		}
	}

	for j := 0; j < n; j++ {
		fmt.Scan(&nval)

		if (b+nval) >= s && (b+nval) <= t {
			ncount++
		}
	}

	fmt.Println(mcount)
	fmt.Println(ncount)
}
