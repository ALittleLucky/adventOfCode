package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t := time.Now()
	crabs := LoadCrabs("input.txt")
	p1 := alignCrabs(crabs)
	p2 := alignCrabsFuelBurn(crabs)

	fmt.Println("Day 7 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))
}

func alignCrabs(crabs []int) int {
	furthest := 0
	for _, c := range crabs {
		if c > furthest {
			furthest = c
		}
	}

	lowestTotal := math.MaxUint32
	for i := 1; i < furthest; i++ {
		total := 0
		for _, c := range crabs {
			if c > i {
				total = total + (c - i)
			} else if c < i {
				total = total + (i - c)
			}
		}
		if lowestTotal > total {
			lowestTotal = total
		}
	}
	return lowestTotal
}
func alignCrabsFuelBurn(crabs []int) int {
	furthest := 0
	for _, c := range crabs {
		if c > furthest {
			furthest = c
		}
	}

	lowestTotal := math.MaxUint32
	for i := 1; i <= furthest; i++ {
		total := 0
		for _, c := range crabs {
			move := 0
			if c > i {
				move = (c - i)
			} else if c < i {
				move = (i - c)
			}
			for j := 1; j <= move; j++ {
				total = total + j
			}

		}
		if lowestTotal > total {
			lowestTotal = total
		}
	}
	return lowestTotal
}
