package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func filter(arr []string, fn func(string) bool) []int {
	filtered := []int{}
	for _, v := range arr {
		if fn(v) {
			conv, err := strconv.Atoi(v)
			if err == nil {
				filtered = append(filtered, conv)
			}
		}
	}
	return filtered
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")
	timeInt, _ := strconv.Atoi(time)

	scanner.Scan()
	distance := strings.ReplaceAll(strings.Split(scanner.Text(), ":")[1], " ", "")
	distanceInt, _ := strconv.Atoi(distance)
	possibleWins  := 0
	for chargeTime := 0; chargeTime <= timeInt; chargeTime++ {
		if chargeTime*(timeInt-chargeTime) > distanceInt {
			possibleWins++
		}

	}
	
	println(possibleWins)
}
