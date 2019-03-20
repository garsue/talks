package main

import (
	"fmt"
	"time"
)

func main() {
	// 2019-01-31
	t1 := time.Date(2019, time.January, 31, 0, 0, 0, 0, time.Local)
	t2 := t1.AddDate(0,1,0)
	fmt.Println(t2)
}

