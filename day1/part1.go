package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//Part1
	d1 := depths()
	lastDepth := 0
	increases := 0

	for i := 0; i < len(d1); i++ {
		if i != 0 && d1[i] > lastDepth {
			increases++
		}
		lastDepth = d1[i]
	}
	p1 := increases
	//Part 2
	d2 := depths()
	SlidingDepth := 0
	SlidingInc := 0

	for i := 0; i+2 < len(d2); i++ {
		if i != 0 && d2[i]+d2[i+1]+d2[i+2] > SlidingDepth {
			SlidingInc++
		}
		SlidingDepth = d2[i] + d2[i+1] + d2[i+2]
	}

	p2 := SlidingInc
	fmt.Println("Day 1 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))
}
