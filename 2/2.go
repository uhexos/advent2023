package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// totalRedCubes := 0
	// totalBlueCubes := 0
	// totalGreenCubes := 0
	for scanner.Scan() {
		games := strings.Split(scanner.Text(), ":")
		sets := strings.Split(games[1], ";")
		for _, turn := range sets {
			colors := strings.Split(turn,",")
			fmt.Println(colors)
			for _,color := range colors{
				fmt.Println(color)
			}
		}
	}
}
