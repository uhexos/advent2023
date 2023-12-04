package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func getNumberOfWins(line string) int {
	winningNumbers := removeEmptyStrings(strings.Split(strings.Split(strings.Split(line, "|")[0], ":")[1], " "))
	elfsNumbers := removeEmptyStrings(strings.Split(strings.Split(line, "|")[1], " "))

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
	return numeberOfWinnningsInLine

}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scratchCards := []string{}
	lineNumber := 1
	for scanner.Scan() {
		key := strconv.Itoa(lineNumber)
		scratchCards[key] = scanner.Text()
		lineNumber++
	}

	for key, line := range scratchCards {
		numOfWins := getNumberOfWins(line)
		for i := 1; i <= numOfWins; i++ {
	        randomNumber := rand.Intn(1000000)
			unique := "," + strconv.Itoa(randomNumber)

			currentKey := strings.Split(key, "")[0]
			keyToDuplicate, _ := strconv.Atoi(currentKey)
			keyToDuplicate += i
			scratchCards[strconv.Itoa(keyToDuplicate)+unique] = scratchCards[strconv.Itoa(keyToDuplicate)]
		}
		// println(winningNumbers, elfsNumbers, math.Pow(2, float64(numeberOfWinnningsInLine-1.0)))
	}
	total := 0
	for _, v := range scratchCards {

		total += len(v)

	}
	fmt.Println(total * len(scratchCards["2"]))
}
