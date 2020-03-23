package main

import "fmt"

func main() {
	var n, temp, sum int64
	fmt.Scan(&n)                    //INPUT DATA
	for i := int64(0); i < n; i++ { //Loop
		fmt.Scan(&temp)
		sum += temp //Tambahkan hasil scan
	}
	fmt.Println(sum) // Print Hasil Sum
}
