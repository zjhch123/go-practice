package main

import (
	"fmt"
	"strconv"
)

func Abs(a int) int {
	var ret int
	switch b := a; {
	case b < 0:
		ret = -b
	case b >= 0:
		ret = b
	default:
		ret = b
	}
	return ret
}

func Season(month int) string {
	var ret string
	switch month {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		ret = "Q1"
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		ret = "Q2"
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		ret = "Q3"
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 12:
		ret = "Q4"
	default:
		ret = "what???"
	}
	return ret
}

func main() {
	a := 10
	if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a < 5")
	}
	fmt.Println(Abs(-142))

	orig := "123"

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(an)

	switch a {
	case 9, 10:
		fmt.Println("a is 9 or 10")
	case 8:
		fmt.Println("a is 8")
	default:
		fmt.Println("a is other")
	}

	fmt.Println(Season(8))

	for index, char := range "在救护车23" {
		fmt.Printf("%d %c\n", index, char)
	}

	for i := 0; i < 10; i++ {
		var v int
		fmt.Println(v)
		v = 5
	}

}
