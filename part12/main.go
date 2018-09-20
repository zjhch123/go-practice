package main

import (
	"fmt"
)

func getX2AndX3(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func main() {
	fmt.Println(getX2AndX3(3))
}
