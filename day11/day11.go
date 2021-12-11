package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const FlashPoint = 10

func main() {
	t := time.Now()
	s, _ := ioutil.ReadFile("input.txt")

	simulationLength := 100
	p1, _ := getFlashes(GetArray(string(s)), simulationLength)
	p2 := getSync(GetArray(string(s)))
	fmt.Println("Day 11 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))

}

func getSync(i [][]int) int {
	allFlash := len(i) * len(i[0])
	currentFlashes := 0
	simmedDays := 0
	for allFlash != currentFlashes {
		currentFlashes, i = getFlashes(i, 1)
		simmedDays++
	}
	return simmedDays
}

func getFlashes(squids [][]int, simLength int) (int, [][]int) {
	flashCount := 0

	flashed := make([][]uint8, len(squids))
	for i := range flashed {
		flashed[i] = make([]uint8, len(squids[0]))
	}
	for i := 0; i < simLength; i++ {
		for i, l := range squids {
			for j := range l {
				squids[i][j]++
				flashed[i][j] = 0
			}
		}
		new := true
		for new {
			new = false
			for i, l := range squids {
				for j := range l {
					if squids[i][j] >= FlashPoint && flashed[i][j] == 0 {
						new = true
						squids = flashPoint(squids, i, j)
						flashed[i][j] = 1
					}
				}
			}
		}
		for i, l := range squids {
			for j := range l {
				if squids[i][j] >= FlashPoint {
					flashCount++
					squids[i][j] = 0
				}
			}
		}

	}
	return flashCount, squids
}

func flashPoint(squids [][]int, i, j int) [][]int {
	if i != 0 { //
		if j != 0 {
			squids[i-1][j-1]++ // top left
		}
		if j != len(squids[i])-1 {
			squids[i-1][j+1]++ // top right
		}
		squids[i-1][j]++ // up
	}
	if j != 0 {
		squids[i][j-1]++ //left
	}

	if j != len(squids[i])-1 {
		squids[i][j+1]++ //right
	}
	if i != len(squids[i])-1 { //
		if j != len(squids[i])-1 {
			squids[i+1][j+1]++ // bottom right
		}
		if j != 0 {
			squids[i+1][j-1]++ // top right
		}
		squids[i+1][j]++ // down
	}
	return squids
}

func GetArray(s string) [][]int {
	lines := strings.Split(s, "\n")
	ar := make([][]int, 0)
	for _, l := range lines {
		currentline := make([]int, 0)
		for _, ns := range l {
			n, _ := strconv.Atoi(string(ns))
			currentline = append(currentline, n)
		}
		if len(currentline) > 1 {
			ar = append(ar, currentline)
		}
	}
	return ar
}
