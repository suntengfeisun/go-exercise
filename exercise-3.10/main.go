package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func one(s string) {
	fmt.Println(comma(s))
}

func commaByBuf(s string) {
	var buf bytes.Buffer
	n := len(s) % 3
	a := []byte(s[:n])
	for _, v := range a {
		buf.WriteByte(v)
	}
	b := []byte(s[n:])
	for k, v := range b {
		if k%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(v)

	}
	fmt.Println(buf.String())
}

func two(s string) {
	commaByBuf(s)
}

func main() {
	myString := "12345789"
	one(myString)
	two(myString)
}
