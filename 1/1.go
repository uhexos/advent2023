package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type CalibrationValue struct {
	first string
	last  string
}

func main() {
	fmt.Println("Hello, world.")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	totalSnow := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		var calibrationValue CalibrationValue

		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				if calibrationValue.first == "" {
					calibrationValue.first = string(char)
				}
				calibrationValue.last = string(char)
			}
		}
		tens, _ := strconv.Atoi(calibrationValue.first)
		tens *= 10
		units, _ := strconv.Atoi(calibrationValue.last)
		totalSnow += tens + units
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalSnow)
	file.Close()
}
