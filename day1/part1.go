package main

import "fmt"

func main() {
	d := depthsTest()
	lastDepth := 0
	increases := 0
	for i := 0; i < len(d); i++ {
		if i != 0 && d[i] > lastDepth {
			increases++
		}
		lastDepth = d[i]
	}
	fmt.Println(increases)
}
