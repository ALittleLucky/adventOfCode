package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Solve("test.txt")
	//SolveP2("t2.txt")
	//SolveP2("test.txt")
	t := time.Now()
	p1 := Solve("input.txt")
	p2 := SolveP2("input.txt")
	fmt.Println("Day 8 Part1:", p1, "Part2:", p2, "Time Taken:", time.Since(t))
}
func SolveP2(inputFile string) int {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		var nums [10]string
		str235 := make([]string, 0) //2,3,5
		str069 := make([]string, 0) //0,6,9
		rawStrings := strings.Split(scanner.Text(), " ")

		for i, w := range rawStrings {
			rawStrings[i] = sortString(w)
		}

		for _, w := range rawStrings[:10] {
			if len(w) == 2 {
				nums[1] = w
			} else if len(w) == 4 {
				nums[4] = w
			} else if len(w) == 3 {
				nums[7] = w
			} else if len(w) == 7 {
				nums[8] = w
			} else if len(w) == 5 {
				str235 = append(str235, w)
			} else {
				str069 = append(str069, w)
			}
			//need 0, 2,3,5,6,9
		}

		for i := range str069 {
			found := 0
			for _, l := range str069[i] {
				if strings.ContainsRune(nums[1], l) {
					found++
				}
			}
			if found == 1 {
				nums[6] = str069[i]
				str069 = append(str069[:i], str069[i+1:]...)
				break // found 6
			}
		}

		for i := range str235 {
			missing := 0

			for _, l := range str235[i] {
				if !strings.ContainsRune(nums[6], l) {
					missing++
				}
			}
			if missing == 0 {
				nums[5] = str235[i]
				str235 = append(str235[:i], str235[i+1:]...)
				break // found 5
			}
		}

		for i := range str069 {
			missing := 0

			for _, l := range str069[i] {
				if !strings.ContainsRune(nums[5], l) {
					missing++
				}
			}
			if missing == 1 {
				nums[9] = str069[i]
				str069 = append(str069[:i], str069[i+1:]...)

				nums[0] = str069[0]
				str069 = nil
				break // found 9 + 0
			}
		}
		for i := range str235 {
			missing := 0

			for _, l := range str235[i] {
				if !strings.ContainsRune(nums[9], l) {
					missing++
				}
			}
			if missing == 0 {
				nums[3] = str235[i]
				str235 = append(str235[:i], str235[i+1:]...)

				nums[2] = str235[0]
				break // found 3,2
			}
		}
		out := ""
		for _, w := range rawStrings[11:] {
			for i, n := range nums {
				if w == n {
					s := strconv.Itoa(i)
					out = out + s
				}
			}
		}
		v, _ := strconv.Atoi(out)
		count += v
	}
	return count
}

func sortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Solve(inputFile string) int {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {

		rawStrings := strings.Split(scanner.Text(), " ")
		for _, w := range rawStrings[11:] {
			if len(w) != 5 && len(w) != 6 {

				c++
			}
		}
	}
	return c

}

// if rune(nums[1][0]) == grid[2] {
// 	grid[1] = rune(nums[1][1])
// } else {
// 	grid[1] = rune(nums[1][0])
// }
// for _, i := range nums[7] {
// 	if i != grid[1] && i != grid[2] {
// 		i = grid[0]
// 		break
// 	}
// }
// for _, i := range nums[4] {
// 	if i != grid[1] && i != grid[2] {
// 		i = grid[0]
// 		break
// 	}
// }
