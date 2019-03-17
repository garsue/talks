package main

import (
	"fmt"
)

func main() {
	var n *int = nil
	fmt.Println("n == nil:", n == nil)

	var i interface{} = n
	fmt.Println("i == n:", i == n)
	fmt.Println("i == nil:", i == nil)
}
