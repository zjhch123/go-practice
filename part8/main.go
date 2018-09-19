package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "Hello, World!"
	fmt.Println(str1)
	str2 := str1 + "zjh"
	fmt.Println(str2)
	fmt.Println(strings.Contains(str2, str1))
	fmt.Println(strings.Index(str2, str1))

	fmt.Println(strings.Replace(str2, "zjh", "zjhch123", -1))

	fmt.Println(strings.Count(str2, "o"))
	fmt.Println(strings.Repeat("zjh\\", 8))
	fmt.Println(strings.ToUpper("zjh"))
	fmt.Println(strings.ToLower("ZJH"))

	fmt.Println(strings.Trim("   zjh   ", " "))

	fmt.Println(strings.Split("zjhch123", "h"))

	fmt.Println(strings.Join(strings.Split("4,1,6,8,3,6,5,0,8", ","), "-"))

	fmt.Println(strconv.Itoa(12346))
	fmt.Println(strconv.FormatFloat(12346.789, 'f', 3, 64)) // 数字, 'b/e/f/g', 精度, 32/64

	fmt.Printf("操作系统是%d位的\n", strconv.IntSize)

	nnn, _ := strconv.Atoi("123456")
	fmt.Println(nnn + 999)

}
