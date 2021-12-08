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
	fish := LoadFish("data.txt")
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
func LoadFish(inputFile string) []int {

	f, _ := ioutil.ReadFile(inputFile)

	intArray := make([]int, 9)
	rawArray := strings.Split(string(f), ",")

	for _, n := range rawArray {
		i, _ := strconv.Atoi(n)
		intArray[i]++
	}

	return intArray
}
