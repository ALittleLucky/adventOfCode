package main

import "fmt"

func main() {
	fish := LoadFish("data.txt")
	simulationLength := 80
	fmt.Println("problem 1:", simulate(fish, simulationLength))

	simulationLength = 256
	fmt.Println("problem 2:", simulate(fish, simulationLength))
}

func simulate(fish []int, l int) int {
	for s := 0; s < l; s++ {
		fish[7] = fish[7] + fish[0]      //add new fish
		fish = append(fish[1:], fish[0]) //move array one down
	}

	count := 0
	for _, f := range fish {
		count += f
	}
	return count
}
