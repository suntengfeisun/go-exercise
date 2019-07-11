package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		if k != 0 {
			fmt.Println(k, v)
		}
	}
}
