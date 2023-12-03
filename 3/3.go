package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getPossibleLocations(arr []string, x int, y int) [][]int {
	var possible [][]int = [][]int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i+x == x && j+y == y) && i+x >= 0 && j+y >= 0 && i+x < len(arr[x]) && j+y < len(arr) {
				possible = append(possible, []int{i + x, j + y})
			}
		}
	}
	return possible
}

func anyValid(arrayToSearch []string, locations [][]int) bool {
	invalids := "0123456789."
	for _, coordinates := range locations {
		if !strings.Contains(invalids, string(arrayToSearch[coordinates[0]][coordinates[1]])) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var arr []string = []string{}
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	for lineNumber, line := range arr {
		var allPossible [][]int = [][]int{}
		num := ""
		for index, value := range line {
			if unicode.IsDigit(value) {
				num += string(value)
				allPossible = append(allPossible, getPossibleLocations(arr, lineNumber, index)...)
			} else  {
				a := anyValid(arr, allPossible)
				if a {
					val, err := strconv.Atoi(num)
					if err == nil {
						total += val
						fmt.Println(lineNumber, index, num)
					}
				}
				// println(num, a)
				num = ""
				allPossible = [][]int{}
			}

			if index == len(line)-1 {
				a := anyValid(arr, allPossible)
				if a {
					val, err := strconv.Atoi(num)
					if err == nil {
						total += val
						fmt.Println(lineNumber, index, num)
					}
				}
				// println(num, a)
				num = ""
				allPossible = [][]int{}
			}
			// println(allPossible)

		}
	}
	fmt.Println(total)
}
