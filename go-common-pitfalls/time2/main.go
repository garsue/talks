package main

import (
	"fmt"
	"time"
)

func main() {
	// 2019-01-31
	t1 := time.Date(2019, time.January, 31, 0, 0, 0, 0, time.Local)
	year, month, _ := t1.Date()
	// 2019-01-01
	t2 := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	t3 := t2.AddDate(0, 1, 0)
	fmt.Println(t3)
}
