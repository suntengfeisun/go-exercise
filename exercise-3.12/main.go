package main

import (
	"fmt"
	"reflect"
	"sort"
)

func animalCompare(one, two string) {
	var oneIntSlice []int
	var twoIntSlice []int
	oneByteSlice := []byte(one)
	twoByteSlice := []byte(two)
	for _, v := range oneByteSlice {
		oneIntSlice = append(oneIntSlice, int(v))
	}
	for _, v := range twoByteSlice {
		twoIntSlice = append(twoIntSlice, int(v))
	}
	sort.Ints(oneIntSlice[:])
	sort.Ints(twoIntSlice[:])
	if reflect.DeepEqual(oneIntSlice, twoIntSlice) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func main() {
	dog := "aasdafj"
	bird := "bbbbb"
	cat := "awaaw"
	bear := "aaaww"
	animalCompare(dog, bear)
	animalCompare(bird, bear)
	animalCompare(cat, bear)
}
