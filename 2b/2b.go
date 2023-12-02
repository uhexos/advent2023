package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ":")
		sets := strings.ReplaceAll(games[1], ";", ",")
		colors := strings.Split(sets, ",")
		// fmt.Println(colors)

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, cube := range colors {
			split := strings.Split(cube, " ")
			number, _ := strconv.Atoi(split[1])
			color := split[2]

			if color == "red" {
				maxRed = findMax(number, maxRed)
			} else if color == "green" {
				maxGreen = findMax(number, maxGreen)
			} else if color == "blue" {
				maxBlue = findMax(number, maxBlue)
			}
			// shouldAdd = shouldAdd && ((color == "red" && number <= totalRedCubes) || (color == "green" && number <= totalGreenCubes) || (color == "blue" && number <= totalBlueCubes))
		}
		total += maxGreen * maxBlue * maxRed

	}
	fmt.Println(total)
}
