package main

import "fmt"

const Pi = 3.14159
var n int = 666

type Color int

const beef, two, c = "eat", 2, "veg"

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

const (
	a = iota
	b = iota
)

const (
	RED Color = iota
	ORANGE
	YELLOW
	GREEN
	BLUE
	INDIGO
	VIOLET
)

var a1 int = 3
var b1 bool

func add(a int) int {
	b := 3
	return a + n + b
}

func main() {
	fmt.Println(add(888))
	fmt.Println(a, b)
	fmt.Println(RED)
	fmt.Println(a1, b1)
}