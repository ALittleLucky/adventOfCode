package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	lines := LoadPoints("data.txt")

	p1(RemoveDiag(lines))
	p1(lines)
}
func p1(lines []line) {

	sizeX, sizeY := arraySize(lines)
	field := make([][]int, sizeX)
	for i := range field {
		field[i] = make([]int, sizeY)
	}

	for _, l := range lines {
		points := getPoints(l)
		for _, p := range points {
			field[p.y][p.x]++
		}
	}
	count := 0
	for i := range field {
		for _, v := range field[i] {
			if v > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
	coulors := []color.RGBA{{255, 255, 255, 0xff},
		{255, 0, 0, 0xff},
		{0, 0, 0, 0xff}}

	start := image.Point{0, 0}
	end := image.Point{sizeX, sizeY}
	img := image.NewRGBA(image.Rectangle{start, end})

	for i := range field {
		for j, v := range field[i] {
			if v == 0 {
				img.Set(i, j, coulors[0])
			} else if v == 1 {
				img.Set(i, j, coulors[1])
			} else {
				img.Set(i, j, coulors[2])
			}
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func getPoints(l line) []point {
	var points []point
	points = append(points, l.start)
	change := 1
	if l.end.x == l.start.x { //horizontal
		if l.start.y > l.end.y {
			change = -1
		}
		for i := l.start.y + change; i != l.end.y; i = i + change {
			points = append(points, point{l.end.x, i})
		}

	} else if l.end.y == l.start.y { //vertical
		if l.start.x > l.end.x {
			change = -1
		}
		for i := l.start.x + change; i != l.end.x; i = i + change {
			points = append(points, point{i, l.end.y})
		}
	} else {
		changeX := 1
		changeY := 1
		if l.start.y > l.end.y {
			changeY = -1
		}
		if l.start.x > l.end.x {
			changeX = -1
		}
		xPos := l.start.x
		for i := l.start.y + changeY; i != l.end.y; i = i + changeY {
			xPos = xPos + changeX
			points = append(points, point{xPos, i})
		}
	}

	points = append(points, l.end)
	return points
}

func RemoveDiag(lines []line) []line {
	var newlines []line
	for i := range lines {
		if lines[i].end.x == lines[i].start.x || lines[i].end.y == lines[i].start.y {
			newlines = append(newlines, lines[i])
		}

	}
	return newlines
}

func arraySize(lines []line) (X int, Y int) {

	var sizeX, sizeY = 0, 0
	for i := range lines {
		if lines[i].start.x > sizeX {
			sizeX = lines[i].start.x
		}
		if lines[i].end.x > sizeX {
			sizeX = lines[i].end.x
		}
		if lines[i].start.y > sizeY {
			sizeY = lines[i].start.y
		}
		if lines[i].end.y > sizeY {
			sizeY = lines[i].end.y
		}
	}
	return sizeX + 1, sizeY + 1
}
