package vtime_test

import (
	"fmt"

	"v8.run/go/vtime"
)

func ExampleNow() {
	fmt.Println(vtime.Now())
	fmt.Println(vtime.VTime[vtime.WallClock]())
	fmt.Println(vtime.UTC[vtime.WallClock]())
}
