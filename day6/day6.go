package main

import (
	"fmt"
	"time"
)

func main() {
	var p1, p2 int
	t := time.Now()
	fish := LoadFish()
	simulationLength := 80
	p1, fish = simulate(fish, simulationLength)
	simulationLength = 256 - simulationLength
	p2, fish = simulate(fish, simulationLength)
	fmt.Println("problem 1:", p1)
	fmt.Println("problem 2:", p2)
	fmt.Println(time.Since(t))
}

func simulate(fish []int, l int) (a int, fishs []int) {
	for s := 0; s < l; s++ {
		fish[7] = fish[7] + fish[0]      //add new fish
		fish = append(fish[1:], fish[0]) //move array one down
	}

	count := 0
	for _, f := range fish {
		count += f
	}
	return count, fish
}
func LoadFish() []int {

	rawArray := []int{2, 1, 2, 1, 5, 1, 5, 1, 2, 2, 1, 1, 5, 1, 4, 4, 4, 3, 1, 2, 2, 3, 4, 1, 1, 5, 1, 1, 4, 2, 5, 5, 5, 1, 1, 4, 5, 4, 1, 1, 4, 2, 1, 4, 1, 2, 2, 5, 1, 1, 5, 1, 1, 3, 4, 4, 1, 2, 3, 1, 5, 5, 4, 1, 4, 1, 2, 1, 5, 1, 1, 1, 3, 4, 1, 1, 5, 1, 5, 1, 1, 5, 1, 1, 4, 3, 2, 4, 1, 4, 1, 5, 3, 3, 1, 5, 1, 3, 1, 1, 4, 1, 4, 5, 2, 3, 1, 1, 1, 1, 3, 1, 2, 1, 5, 1, 1, 5, 1, 1, 1, 1, 4, 1, 4, 3, 1, 5, 1, 1, 5, 4, 4, 2, 1, 4, 5, 1, 1, 3, 3, 1, 1, 4, 2, 5, 5, 2, 4, 1, 4, 5, 4, 5, 3, 1, 4, 1, 5, 2, 4, 5, 3, 1, 3, 2, 4, 5, 4, 4, 1, 5, 1, 5, 1, 2, 2, 1, 4, 1, 1, 4, 2, 2, 2, 4, 1, 1, 5, 3, 1, 1, 5, 4, 4, 1, 5, 1, 3, 1, 3, 2, 2, 1, 1, 4, 1, 4, 1, 2, 2, 1, 1, 3, 5, 1, 2, 1, 3, 1, 4, 5, 1, 3, 4, 1, 1, 1, 1, 4, 3, 3, 4, 5, 1, 1, 1, 1, 1, 2, 4, 5, 3, 4, 2, 1, 1, 1, 3, 3, 1, 4, 1, 1, 4, 2, 1, 5, 1, 1, 2, 3, 4, 2, 5, 1, 1, 1, 5, 1, 1, 4, 1, 2, 4, 1, 1, 2, 4, 3, 4, 2, 3, 1, 1, 2, 1, 5, 4, 2, 3, 5, 1, 2, 3, 1, 2, 2, 1, 4}
	intArray := make([]int, 9)

	for _, n := range rawArray {

		intArray[n]++
	}

	return intArray
}
