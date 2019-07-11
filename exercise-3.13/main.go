package main

import (
	"fmt"
	"math"
)

// 常量只能用常量赋值，不能用变量
// const (
// 	IB = iota
// 	KB = math.Pow(1000, iota)
// )
const (
	IB = 1
	KB = IB * 1000
)

func main() {
	fmt.Println(IB)
	fmt.Println(KB)
	fmt.Println(math.Pow(1000, 2))

}
