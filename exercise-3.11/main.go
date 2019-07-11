package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

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

func one(s string) {
	commaByBuf(s)
}

func commaByBufNew(s string) {
	i, _ := strconv.ParseFloat(s, 64)

	var buf bytes.Buffer
	var n int
	var m int
	c := 0
	if i < 0 {
		buf.WriteByte('-')
		s = string(s[1:])
	}
	if strings.LastIndex(s, ".") != -1 {
		m++
	}
	n = (len(s) - m) % 3
	a := []byte(s[:n])
	for _, v := range a {
		buf.WriteByte(v)
	}
	b := []byte(s[n:])
	for _, v := range b {
		if v == '.' {
			buf.WriteByte(v)
			continue
		}
		if c%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(v)
		c++
	}
	fmt.Println(buf.String())
}

func two(s string) {
	commaByBufNew(s)
}

func main() {
	myString := "-1234.5789"
	one(myString)
	two(myString)
}
