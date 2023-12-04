package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	winnings := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		winningNumbers := removeEmptyStrings(strings.Split(strings.Split(strings.Split(scanner.Text(), "|")[0], ":")[1], " "))
		elfsNumbers := removeEmptyStrings(strings.Split(strings.Split(scanner.Text(), "|")[1], " "))

		dict := make(map[string]int)

		for _, v := range winningNumbers {
			dict[v] = 0
		}

		for _, v := range elfsNumbers {
			if _, ok := dict[v]; ok {
				dict[v] += 1
			}
		}
		numeberOfWinnningsInLine := 0
		for _, v := range dict {
			if v != 0 {
				numeberOfWinnningsInLine += v
			}
		}
		if numeberOfWinnningsInLine > 0 {
			winnings += math.Pow(2, float64(numeberOfWinnningsInLine-1.0))

		}
		// println(winningNumbers, elfsNumbers, math.Pow(2, float64(numeberOfWinnningsInLine-1.0)))

	}
	fmt.Println(winnings)
}
