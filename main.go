package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	thismorning = flag.Bool("thismorning", false, "midnight this morning (local time)")
)

func main() {
	flag.Parse()

	var t time.Time
	switch {
	case *thismorning:
		t = time.Now()
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	default:
		t = time.Now()
	}
	fmt.Println(t.Unix())
}
