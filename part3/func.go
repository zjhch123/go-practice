package main

import "fmt"

type (
	IZ int
	FZ float64
	STR string
)

func Sum(a IZ, b IZ) IZ {
	return a + b
}

func SumWithArgs(a int, b int) (int, int, int) {
	return a, b, a + b
}

func main() {
	fmt.Println(Sum(2, 3))
	fmt.Println(SumWithArgs(2, 3))
}