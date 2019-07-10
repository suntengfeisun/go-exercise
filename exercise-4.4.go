package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int) {
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
}

func one(myArray [5]int) {
	mySlice := myArray[:]
	rotate(mySlice)
	fmt.Println("one", myArray)
}

func rotateNew(s []int) (sNew []int) {
	n := 0
	temp := make([]int, len(s)*2)
	for i, v := range s {
		if i >= 2 {
			temp[i] = v
		} else {
			temp[len(s)+i] = v
		}
		n++
	}
	sNew = temp[2 : len(s)+2]
	return
}

func two(myArray [5]int) {
	mySlice := myArray[:]
	myArrayNew := rotateNew(mySlice)
	fmt.Println("two", myArray)
	fmt.Println("two", myArrayNew)
}

func main() {
	var myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	// fist of all, let's see how to use the origin reverse function
	one(myArray)
	// my new reverse
	two(myArray)
}
