package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println("Day 4 Part1:", p1(), "Part2:", p2(), "Time Taken:", time.Since(t))

}
func p1() int {
	nums, cards := LoadGame("actual.txt")
	winner, num := findWinner(nums, cards)
	return countUnmarked(winner) * num
}

func p2() int {
	nums, cards := LoadGame("actual.txt")

	for len(cards) > 0 {
		winner, num := findWinner(nums, cards)
		if len(cards) == 1 {
			return countUnmarked(winner) * num
		}
		index := findCard(cards, winner.index)
		cards = append(cards[:index], cards[index+1:]...)
	}
	return 0
}
func findCard(cards []card, w int) int {
	for i, c := range cards {
		if c.index == w {
			return i
		}
	}
	return 0
}
func countUnmarked(c card) int {
	count := 0
	for i := range c.set[0] {
		for j := range c.set[0] {
			if !c.set[i][j].marked {
				count = count + c.set[i][j].val
			}

		}
	}
	return count
}

func findWinner(Nums []int, cards []card) (c card, n int) {
	for _, n := range Nums {
		for k, c := range cards {
			for i := range c.set {
				for j := range c.set {
					if c.set[i][j].val == n {
						cards[k].set[i][j].marked = true

						if checkWinner(cards[k], i, j) {
							return cards[k], n
						}

					}
				}
			}

		}
	}

	return c, 0
}

func checkWinner(c card, x int, y int) bool {

	for i := range c.set[0] {
		if !c.set[x][i].marked {
			break
		}
		if i == 4 {
			return true
		}
	}
	for i := range c.set[0] {
		if !c.set[i][y].marked {
			break
		}
		if i == 4 {
			return true
		}
	}
	return false
}
