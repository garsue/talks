package main

import (
	"fmt"
)

func main() {
	var x int
	defer fmt.Println(x)
	x = 1
}
