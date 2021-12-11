package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//Part1

	t := time.Now()
	d1 := data()
	var gammaRate []int = d1[0]

	for i := 1; i < len(d1); i++ {
		for j := 0; j < len(d1[i]); j++ {
			if d1[i][j] == 0 {
				gammaRate[j]--
			} else {
				gammaRate[j]++
			}
		}
	}
	GammaOutput := 0
	EpsilonOutput := 0
	for i := len(gammaRate) - 1; i > -1; i-- {

		if gammaRate[i] > 0 {
			GammaOutput = GammaOutput + int(math.Pow(2, float64(len(gammaRate)-i-1)))
		} else {
			EpsilonOutput = EpsilonOutput + int(math.Pow(2, float64(len(gammaRate)-i-1)))
		}
	}
	p1 := GammaOutput * EpsilonOutput
	//part2
	p2 := oxygen() * co2()

	fmt.Println("Day 3 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))

}

func oxygen() int {
	var list0 [][]int
	var list1 [][]int
	currentList := data()
	currentPos := 0
	for j := 0; j < len(currentList[0]); j++ {
		for i := 0; i < len(currentList); i++ {
			if currentList[i][currentPos] == 0 {
				list0 = append(list0, currentList[i])
			} else {
				list1 = append(list1, currentList[i])
			}

		}
		if len(list0) > len(list1) {
			currentList = list0
		} else {
			currentList = list1
		}
		list0 = nil
		list1 = nil
		currentPos++
		if len(currentList) == 1 {
			j = math.MaxInt32
		}
	}

	Output := 0
	for i := len(currentList[0]) - 1; i > -1; i-- {

		if currentList[0][i] > 0 {
			Output = Output + int(math.Pow(2, float64(len(currentList[0])-i-1)))
		}
	}
	return Output
}

func co2() int {
	var list0 [][]int
	var list1 [][]int
	currentList := data()
	currentPos := 0
	for j := 0; j < len(currentList[0]); j++ {
		for i := 0; i < len(currentList); i++ {
			if currentList[i][currentPos] == 1 {
				list0 = append(list0, currentList[i])
			} else {
				list1 = append(list1, currentList[i])
			}

		}
		if len(list0) < len(list1) {
			currentList = list0
		} else {
			currentList = list1
		}
		list0 = nil
		list1 = nil
		currentPos++
		if len(currentList) == 1 {
			j = math.MaxInt32
		}
	}
	Output := 0
	for i := len(currentList[0]) - 1; i > -1; i-- {

		if currentList[0][i] > 0 {
			Output = Output + int(math.Pow(2, float64(len(currentList[0])-i-1)))
		}
	}
	return Output
}
