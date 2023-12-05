package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func buildMap(source int, destination int, lenght int) []int {
	maxInt := source + lenght
	arr := make([]int, maxInt)
	count := 0
	for i := range arr {

		if i < source {
			arr[i] = i
		} else {
			arr[i] = destination + count
			count++
		}
	}
	return arr
}
func main() {
	// seedToSoilMap := [][]int{}

	a := buildMap(50, 52, 48)

	m := make(map[string][][]int)

	// Add data to the map
	m["seed-to-soil"] = [][]int{}
	m["soil-to-fertilizer"] = [][]int{}
	m["fertilizer-to-water"] = [][]int{}
	m["water-to-light"] = [][]int{}
	m["light-to-temperature"] = [][]int{}
	m["temperature-to-humidity"] = [][]int{}
	m["humidity-to-location"] = [][]int{}
	fmt.Println(a)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println("for ran")
		if strings.Contains(scanner.Text(), "-") {
			key := strings.Split(scanner.Text(), " ")[0]
			for scanner.Text() != "\n" && scanner.Text() != "" {
				scanner.Scan()
				if scanner.Text() != "" && !strings.Contains(scanner.Text(), "-") {
					substrings := strings.Split(scanner.Text(), " ")
					destination, _ := strconv.Atoi(substrings[0])
					source, _ := strconv.Atoi(substrings[1])
					length, _ := strconv.Atoi(substrings[2])
					m[key] = append(m[key], buildMap(source, destination, length))
				}

			}
		}
	}
	fmt.Println(m)
}
