package main

import (
	"fmt"
	"os"
)

func one() {
	for k, v := range os.Args {
		if k != 0 {
			fmt.Println(k, v)
		}
	}
}

func two() {
	for k, v := range os.Args[1:] {
		fmt.Println(k, v)
	}
}

// two ways have different index values
func main() {
	one()
	two()
}
