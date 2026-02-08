package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Format("2006-01-02 Monday 15:04:05"))

	create := time.Date(2026, time.April, 20, 12, 12,0, 0, time.UTC)
	fmt.Println(create.Format("2006-01-02 Monday 15:04:05"))
}