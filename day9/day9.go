package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int
	y int
	v int
}

func main() {
	t := time.Now()
	s, _ := ioutil.ReadFile("input.txt")
	ar, lp := GetArrayAndLowPoints(string(s))
	p1 := getRisk(lp)
	p2 := getBiggestBasins(ar, lp)
	fmt.Println("Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))

}

func getBiggestBasins(ar [][]int, points []point) int {
	basins := make([]int, 0)
	var biggest [3]int
	for _, p := range points {
		basins = append(basins, getBasin(ar, p))
	}
	//get biggest basins
	for _, b := range basins {
		t1 := b
		for i := range biggest {
			if t1 > biggest[i] {
				t2 := biggest[i]
				biggest[i] = t1
				t1 = t2

			}
		}
	}
	c := 1
	for _, b := range biggest {
		c *= b
	}
	return c
}
func getBasin(ar [][]int, p point) int {
	basin := make([]point, 0)
	basin = append(basin, p)
	newpoints := make([]point, 0)
	processingpoints := make([]point, 0)
	newpoints = append(newpoints, p)
	for len(newpoints) > 0 {
		processingpoints = nil
		for _, i := range newpoints {
			ptpoints, b := getFlowPoints(ar, i, basin)
			basin = b
			processingpoints = append(processingpoints, ptpoints...)
		}
		newpoints = processingpoints

	}
	return len(basin)
}

func getFlowPoints(ar [][]int, p point, basin []point) (newpoints []point, allpoints []point) {
	newpoints = make([]point, 0)
	if p.x != 0 {
		v := ar[p.x-1][p.y]
		if v > p.v && v != 9 && !pointExists(p.x-1, p.y, basin) {
			newpoints = append(newpoints, point{p.x - 1, p.y, v})
		}
	}
	if p.x != len(ar)-1 {
		v := ar[p.x+1][p.y]
		if v > p.v && v != 9 && !pointExists(p.x+1, p.y, basin) {
			newpoints = append(newpoints, point{p.x + 1, p.y, v})
		}
	}
	if p.y != 0 {
		v := ar[p.x][p.y-1]
		if v > p.v && v != 9 && !pointExists(p.x, p.y-1, basin) {
			newpoints = append(newpoints, point{p.x, p.y - 1, v})
		}
	}
	if p.y != len(ar[p.x])-1 {
		v := ar[p.x][p.y+1]
		if v > p.v && v != 9 && !pointExists(p.x, p.y+1, basin) {
			newpoints = append(newpoints, point{p.x, p.y + 1, v})
		}
	}
	basin = append(basin, newpoints...)

	return newpoints, basin
}
func getRisk(lp []point) int {
	r := 0
	for _, n := range lp {
		r += n.v + 1
	}
	return r
}

func GetArrayAndLowPoints(s string) ([][]int, []point) {
	lines := strings.Split(s, "\n")
	ar := make([][]int, 0)
	points := make([]point, 0)
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

	for i, l := range ar {
		for j, ns := range l {
			up, down, left, right := 100, 100, 100, 100
			if i != 0 {
				up = ar[i-1][j]
			}
			if i != len(ar)-1 {
				down = ar[i+1][j]
			}
			if j != 0 {
				left = ar[i][j-1]
			}
			if j != len(l)-1 {
				right = ar[i][j+1]
			}
			if ns < up && ns < down && ns < left && ns < right {
				points = append(points, point{i, j, ns})
			}

		}

	}
	return ar, points
}

func pointExists(x int, y int, points []point) bool {
	for _, p := range points {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}
