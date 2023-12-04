package main

import (
	"bufio"
	"fmt"
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
	// winnings := 0.0
	scanner := bufio.NewScanner(file)
	scratchCards := make(map[int][]string)
	lineNumber := 1
	for scanner.Scan() {
		scratchCards[lineNumber] = append(scratchCards[lineNumber], scanner.Text())
		lineNumber++
	}

	for key,line:= range scratchCards{
		numOfWins := getNumberOfWins(line[0])
		
		for i := 1; i <= numOfWins; i++ {
			scratchCards[key+i] = append(scratchCards[key+i], scratchCards[key+i][0])
		}
		println(line[0],numOfWins)
	}
	total:=0
	for _,v:=range scratchCards{
		for _,j:=range v{
			total += len(j)
		}
	}
	fmt.Println(total * len(scratchCards[2]))
}
