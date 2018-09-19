package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {
	a, b := 4, 1
	b, a = a, b
	fmt.Println(a, b)
	goos := runtime.GOOS
	fmt.Printf("The operating system is %v\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}