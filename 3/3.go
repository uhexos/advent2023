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
			//go round index basically but make sure you dont go out of bounds so prevent negatives and too large indexes, also dont include itself
			//contains duplicates but go doesnt have sets and I am too lazy to figure it out using a map
		
			if !(i+x == x && j+y == y) && i+x >= 0 && j+y >= 0 && i+x < len(arr[x]) && j+y < len(arr) {
				possible = append(possible, []int{i + x, j + y})
			}
		}
	}
	return possible
}

func anyValid(arrayToSearch []string, locations [][]int) (bool, []int) {
	invalids := "0123456789."
	for _, coordinates := range locations {
		if !strings.Contains(invalids, string(arrayToSearch[coordinates[0]][coordinates[1]])) {
			return true, []int{coordinates[0], coordinates[1]}
		}
	}
	return false, nil
}

func checkAndAdd(arr []string, allPossible [][]int, num string, total *int) []int {
	a, symbolCoordinates := anyValid(arr, allPossible)
	if a {
		val, err := strconv.Atoi(num)
		if err == nil {
			*total += val
		}
	}
	return symbolCoordinates
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var arr []string = []string{}
	total := 0
	gearRatio := 0
	allValidSymbolLocations := make(map[string][]string)
	//build entire array first otherwise you run into index out of bound exceptions cause a neighbour might be on the lower line 
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
			} else {
				symbolLocation := checkAndAdd(arr, allPossible, num, &total)
				// use coordinates as a key so you can group numbers that share a * symbol, condition duplicated below 
				// any change here must happen below as well. TODO 🤣🤣🤣🤣🤣🤣🤣🤣 move logic to single function
				if symbolLocation != nil && arr[symbolLocation[0]][symbolLocation[1]] == '*' {
					key := strconv.Itoa(symbolLocation[0])+","+strconv.Itoa(symbolLocation[1])
					allValidSymbolLocations[key] = append(allValidSymbolLocations[key], num)
				}
				num = ""
				allPossible = [][]int{}
			}

			//handle case where the number ends the string, without this condition parts are omitted
			if index == len(line)-1 {
				symbolLocation := checkAndAdd(arr, allPossible, num, &total)
				if symbolLocation != nil && arr[symbolLocation[0]][symbolLocation[1]] == '*' {
					key := strconv.Itoa(symbolLocation[0])+","+strconv.Itoa(symbolLocation[1])
					allValidSymbolLocations[key] = append(allValidSymbolLocations[key], num)
				}
				num = ""
				allPossible = [][]int{}
			}
		}
	}

	//for part 2: find all valid symbol locations that have two integers only then multiply and sum
	for _, val := range allValidSymbolLocations {
		if len(val) == 2 {
			val1, _ := strconv.Atoi(val[0])
			val2, _ := strconv.Atoi(val[1])

			gearRatio += val1 * val2
		}
	}
	fmt.Println(total, gearRatio)
}
