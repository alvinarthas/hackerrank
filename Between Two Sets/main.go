package main

import "fmt"

func main() {
	var n, m, x, c int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	b := make([]int, m)

	// Terima inputan deret pertama
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	// Terima inputan deret kedua
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	// lakukan loop selama nilai x bertambah dan nilai x <= dari array pertama deret kedua
	for x <= b[0] {
		factor := true

		x += a[0]

		//cek deret pertama
		for j := 0; j < n; j++ {
			if x%a[j] != 0 { //cek apakah x habis membagi deret pertama
				factor = false
				break
			}
		}
		// cek deret kedua
		if factor {
			for j := 0; j < m; j++ {
				if b[j]%x != 0 { // cek apakah deret kedua habis membagi x
					factor = false
					break
				}
			}
		}
		//update counter
		if factor {
			c++
		}
	}

	fmt.Println(c)

}
