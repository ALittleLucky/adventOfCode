package main

import "fmt"

func main() {
	//Part1
	m1 := moves()
	x1, y1 := 0, 0

	for i := 0; i < len(m1); i++ {
		switch m1[i].direction {
		case "forward":
			x1 = x1 + m1[i].value
		case "down":
			y1 = y1 + m1[i].value
		case "up":
			y1 = y1 - m1[i].value
		}
		fmt.Println("x:", x1, "y:", y1)
	}

	fmt.Println("output:", x1*y1)

	//Part2
	m2 := moves()
	x2, y2, aim := 0, 0, 0

	for i := 0; i < len(m2); i++ {
		switch m2[i].direction {
		case "forward":
			x2 = x2 + m2[i].value
			y2 = y2 + aim*m2[i].value
		case "down":
			aim = aim + m2[i].value
		case "up":
			aim = aim - m2[i].value
		}
		fmt.Println("x:", x2, "y:", y2)
	}

	fmt.Println("output:", x2*y2)
}
