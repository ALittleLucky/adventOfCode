package main

import "fmt"

func main() {
	fish := LoadFish("data.txt")
	simulationLength := 80
	fmt.Println("problem 1:", simulate(fish, simulationLength))

	simulationLength = 256
	fmt.Println("problem 2:", simulate(fish, simulationLength))
}

func simulate(fish [9]int, l int) int {
	var newfish [9]int
	for s := 0; s < l; s++ {
		fish[7] = fish[7] + fish[0]
		newfish[8] = fish[0]
		for i := len(fish) - 2; i >= 0; i-- {
			newfish[i] = fish[i+1]
		}
		fish = newfish
	}
	count := 0

	for i := range fish {
		count = count + fish[i]
	}
	return count
}
