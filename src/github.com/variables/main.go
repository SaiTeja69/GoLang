package main

import (
	"fmt"
	"strconv"
)

var (
	x int = 1
	y int = 2
)

func main() {
	var i int = 54
	a := 45
	var j float32
	j = 22
	j = float32(i)
	str := "sai"
	l := strconv.Itoa(a)
	fmt.Println(i, a, j, str, x, l)
	fmt.Printf("%v, %T", j, j)
}
