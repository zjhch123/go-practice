package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d ", a)
	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		a := rand.Intn(10)
		fmt.Printf("%d ", a)
	}
}
