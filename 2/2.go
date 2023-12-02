package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	totalRedCubes := 12
	totalBlueCubes := 14
	totalGreenCubes := 13
	total := 0
	count := 1

	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ":")
		sets := strings.ReplaceAll(games[1], ";", ",")
		colors := strings.Split(sets, ",")
		// fmt.Println(colors)
		shouldAdd := true
		for _, cube := range colors {
			split := strings.Split(cube, " ")
			number, _ := strconv.Atoi(split[1])
			color := split[2]

			shouldAdd = shouldAdd && ((color == "red" && number <= totalRedCubes) || (color == "green" && number <= totalGreenCubes) || (color == "blue" && number <= totalBlueCubes))

		}
		if shouldAdd {
			total += count
		}
		count += 1
	}
	fmt.Println(total)
}
