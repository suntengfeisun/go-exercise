package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func one(myArray [3]int) {
	mySlice := myArray[:]
	reverse(mySlice)
	fmt.Println("one", myArray)
}

func reverseNew(p *[3]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		// p[i], p[j] = p[j], p[i]
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

func two(myArray [3]int) {
	myPointer := &myArray
	reverseNew(myPointer)
	fmt.Println("two", myArray)
}

// 反转数组
// to reverse an array
func main() {
	var myArray = [3]int{1, 2, 3}
	fmt.Println(myArray)
	// fist of all, let's see how to use the origin reverse function
	one(myArray)
	// my new reverse
	two(myArray)
}
