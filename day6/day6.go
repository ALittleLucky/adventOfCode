package main

import "fmt"

func main() {
	fish := LoadFish("data.txt")
	simulationLength := 80
	fmt.Println("problem 1", simulate(fish, simulationLength))

	simulationLength = 256
	fmt.Println("problem 2", simulate(fish, simulationLength))
}

func simulate(fish [9]int, l int) int {
	newfish := fish
	for s := 0; s < l; s++ {

		for i := len(fish) - 1; i >= 0; i-- {
			if i == 8 {
				newfish[i] = fish[0] // add new fish
			} else if i == 0 {
				newfish[6] = fish[i] + fish[7] // loop fish
				newfish[0] = fish[i+1]
			} else {
				newfish[i] = fish[i+1]
			}
		}

		fish = newfish
	}
	count := 0

	for i := range fish {
		count = count + fish[i]
	}
	return count
}
