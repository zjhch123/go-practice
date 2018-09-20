package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	var i2 = 10
	fmt.Printf("An integer: %d, it's location in memory: %v\n", i1, &i1)
	var p *int = &i1

	fmt.Println(p, *p)
	fmt.Println(i1, i2)
	i2 = *p
	fmt.Println(i2, &i2, i2 == i1, &i2 == &i1)
	fmt.Println(i2 == *(&i2))

	s := "hello world"
	var ptr *string = &s
	fmt.Println(ptr)
	*ptr = "giao"
	fmt.Println(ptr)
	fmt.Println(s)

}
