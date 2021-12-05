package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	start, end point
}
type point struct {
	x int
	y int
}

func LoadPoints(inputFile string) (lines []line) {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineString []string
		lineString = strings.Split(scanner.Text(), " ")
		var l line
		l.start = pointFromString(lineString[0])
		l.end = pointFromString(lineString[2])

		lines = append(lines, l)

	}

	return lines
}

func pointFromString(pointString string) point {
	var x, y int

	ar := strings.Split(pointString, ",")
	x, _ = strconv.Atoi(ar[0])
	y, _ = strconv.Atoi(ar[1])

	return point{x, y}
}
