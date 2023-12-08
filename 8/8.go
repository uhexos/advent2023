package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mapOfNodes := map[string][]string{}

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	searchPattern := scanner.Text()
	scanner.Scan()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " = ")
		leftRight := strings.ReplaceAll(line[1], "(", "")
		leftRight = strings.ReplaceAll(leftRight, ")", "")
		leftRight = strings.ReplaceAll(leftRight, " ", "")
		mapOfNodes[line[0]] = []string{strings.Split(leftRight, ",")[0], strings.Split(leftRight, ",")[1]}
		fmt.Println(mapOfNodes)
	}
	current := "AAA"
	index := 0
	for current != "ZZZ" {
		if searchPattern[index%len(searchPattern)] == 'L' {
			current = mapOfNodes[current][0]
		} else {
			current = mapOfNodes[current][1]
		}
		fmt.Println(current)
		index++
	}

	fmt.Println(index)
}
