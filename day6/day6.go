package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	var p1, p2 int
	t := time.Now()
	fish := LoadFish()
	p1 = simulate(fish, 80)
	p2 = simulate(fish, 256)

	fmt.Println("Day 6 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))

}

func simulate(fish [9]int, l int) (a int) {

	for s := 0; s < l; s++ {
		fish[(s+7)%9] += fish[(s)%9] //add new fish
	}

	count := 0
	for _, f := range fish {
		count += f
	}
	return count
}
func LoadFish() [9]int {
	file, _ := ioutil.ReadFile("data.txt")

	rawArray := strings.Split(string(file), ",")
	//rawArray := []int{2, 1, 2, 1, 5, 1, 5, 1, 2, 2, 1, 1, 5, 1, 4, 4, 4, 3, 1, 2, 2, 3, 4, 1, 1, 5, 1, 1, 4, 2, 5, 5, 5, 1, 1, 4, 5, 4, 1, 1, 4, 2, 1, 4, 1, 2, 2, 5, 1, 1, 5, 1, 1, 3, 4, 4, 1, 2, 3, 1, 5, 5, 4, 1, 4, 1, 2, 1, 5, 1, 1, 1, 3, 4, 1, 1, 5, 1, 5, 1, 1, 5, 1, 1, 4, 3, 2, 4, 1, 4, 1, 5, 3, 3, 1, 5, 1, 3, 1, 1, 4, 1, 4, 5, 2, 3, 1, 1, 1, 1, 3, 1, 2, 1, 5, 1, 1, 5, 1, 1, 1, 1, 4, 1, 4, 3, 1, 5, 1, 1, 5, 4, 4, 2, 1, 4, 5, 1, 1, 3, 3, 1, 1, 4, 2, 5, 5, 2, 4, 1, 4, 5, 4, 5, 3, 1, 4, 1, 5, 2, 4, 5, 3, 1, 3, 2, 4, 5, 4, 4, 1, 5, 1, 5, 1, 2, 2, 1, 4, 1, 1, 4, 2, 2, 2, 4, 1, 1, 5, 3, 1, 1, 5, 4, 4, 1, 5, 1, 3, 1, 3, 2, 2, 1, 1, 4, 1, 4, 1, 2, 2, 1, 1, 3, 5, 1, 2, 1, 3, 1, 4, 5, 1, 3, 4, 1, 1, 1, 1, 4, 3, 3, 4, 5, 1, 1, 1, 1, 1, 2, 4, 5, 3, 4, 2, 1, 1, 1, 3, 3, 1, 4, 1, 1, 4, 2, 1, 5, 1, 1, 2, 3, 4, 2, 5, 1, 1, 1, 5, 1, 1, 4, 1, 2, 4, 1, 1, 2, 4, 3, 4, 2, 3, 1, 1, 2, 1, 5, 4, 2, 3, 5, 1, 2, 3, 1, 2, 2, 1, 4}
	var intArray [9]int

	for _, n := range rawArray {
		i, _ := strconv.Atoi(n)
		// i :=n
		intArray[i]++
	}

	return intArray
}
