package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scan(&x1, &v1, &x2, &v2)

	if x1 == x2 && v1 == v2 {
		fmt.Println("YES")
	} else if x1 == x2 && v1 != v2 {
		fmt.Println("NO")
	} else {
		if v1 > v2 && ((x2-x1)%(v1-v2)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
