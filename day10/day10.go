package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type pair struct {
	open, close rune
	v           int
	f           int
}
type invalidLine struct {
	s string
	p pair
}
type incorrectLine struct {
	s     string
	pairs []pair
}

func main() {
	t := time.Now()
	s, _ := ioutil.ReadFile("input.txt")
	pairs := getPairs()
	inL, icl := GetInvalidLines(string(s), pairs)
	p1 := 0
	for _, l := range inL {
		p1 += l.p.v
	}
	scores := make([]int, 0)
	for _, i := range icl {
		lineValue := 0
		for _, p := range i.pairs {
			for _, pf := range pairs {
				if p == pf {
					lineValue *= 5
					lineValue += p.f

					break
				}
			}
		}
		scores = append(scores, lineValue)
	}

	sort.Ints(scores)

	p2 := scores[len(scores)/2]
	fmt.Println("Day 10 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))

}

func getPairs() []pair {
	pairs := make([]pair, 0)
	pairs = append(pairs, pair{'(', ')', 3, 1})
	pairs = append(pairs, pair{'[', ']', 57, 2})
	pairs = append(pairs, pair{'{', '}', 1197, 3})
	pairs = append(pairs, pair{'<', '>', 25137, 4})
	return pairs
}

func GetInvalidLines(s string, pairs []pair) (inl []invalidLine, icl []incorrectLine) {
	lines := strings.Split(s, "\n")
	ar := make([]invalidLine, 0)
	incorrectLineStrings := make([]string, 0)
	incorrectLines := make([]incorrectLine, 0)
	for _, l := range lines {
		temp := []rune(l)
		for len(temp) > 0 {

			for i := 0; i < len(temp); i++ {
				t := temp[i]
				for j := 0; j < len(pairs); j++ {
					p := pairs[j]
					if t == p.open {

						if i != len(temp)-1 {
							if temp[i+1] == p.close {

								temp = append(temp[:i], temp[i+2:]...)

								j = -1
								i = j
								break
							}
						} else { //valid line?
							incorrectLineStrings = append(incorrectLineStrings, string(temp))
							temp = nil
						}
					}
					if t == p.close {
						ar = append(ar, invalidLine{l, p})

						temp = nil
						break
					}
				}
			}
		}
	}
	for _, l := range incorrectLineStrings {
		tempP := make([]pair, 0)
		for _, r := range l {
			for _, p := range pairs {
				if r == p.open {
					tempP = append([]pair{p}, tempP...)
				}
			}
		}
		incorrectLines = append(incorrectLines, incorrectLine{l, tempP})
	}
	return ar, incorrectLines
}
