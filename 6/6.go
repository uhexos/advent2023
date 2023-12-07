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
	time := []string{}
	distance := []string{}

	scanner.Scan()
	time = append(time, strings.Split(scanner.Text(), " ")...)
	timeInt := filter(time, func(str string) bool { return str != "" })

	scanner.Scan()
	distance = append(distance, strings.Split(scanner.Text(), " ")...)
	distanceInt := filter(distance, func(str string) bool { return str != "" })
	possibleLosses := make([]int, len(distanceInt))

	for index, raceTIme := range timeInt {
		dist := distanceInt[index]
		for chargeTime := 0; chargeTime <= raceTIme; chargeTime++ {
			if chargeTime*(raceTIme-chargeTime) > dist {
				possibleLosses[index]++
			}
			
		}

	}
	margin :=1 
	for _,v:= range possibleLosses{
		margin *= v
	}
	println(margin)
}
