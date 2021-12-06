package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadFish(inputFile string) (ar [9]int) {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var intArray [9]int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rawArray := strings.Split(scanner.Text(), ",")

	for _, n := range rawArray {
		i, _ := strconv.Atoi(n)
		intArray[i]++
	}

	return intArray
}
