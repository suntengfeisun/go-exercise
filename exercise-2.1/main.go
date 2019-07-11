package main

import (
	"fmt"
	tempconv "go-exercise/tempconv"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToK(0))
	fmt.Println(tempconv.KToC(0))
}
