package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	randomString1 := uuid.New().String()[:8]
	randomString2 := uuid.New().String()[:8]
	clientSecret := fmt.Sprintf("%s.%s", randomString1, randomString2)

	fmt.Println(clientSecret)
}
