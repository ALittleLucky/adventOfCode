package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	cards []card
	Nums  []int
}
type card struct {
	set   [5][5]bingoNumber
	index int
}
type bingoNumber struct {
	val    int
	marked bool
}

func LoadGame(inputFile string) (Nums []int, cards []card) {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//Read nums
	scanner.Scan()
	NumsStrings := strings.Split(scanner.Text(), ",")

	for _, i := range NumsStrings {
		n, e := strconv.Atoi(i)
		if e != nil {
			break
		}
		Nums = append(Nums, n)
	}

	for scanner.Scan() {
		var lineString [5][]string
		for i := range lineString {
			scanner.Scan()
			lineString[i] = strings.Split(scanner.Text(), " ")
		}

		var c [5][5]bingoNumber
		for i, s := range lineString {
			pos := 0
			for _, v := range s {

				n, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				c[i][pos] = bingoNumber{n, false}
				pos++
			}
		}
		var crd card
		crd.set = c
		crd.index = len(cards)
		cards = append(cards, crd)
	}
	return Nums, cards
}
