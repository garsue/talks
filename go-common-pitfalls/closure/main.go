package main

import (
	"fmt"
)

func main() {
	var fs []func()
	for i := 0; i < 3; i++ {
		fs = append(fs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range fs {
		f()
	}
}
