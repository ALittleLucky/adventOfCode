package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadCrabs(inputFile string) []int {

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var intArray []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rawArray := strings.Split(scanner.Text(), ",")

	for _, n := range rawArray {
		i, _ := strconv.Atoi(n)
		intArray = append(intArray, i)
	}

	return intArray
}
