package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// type CalibrationValue struct {
// 	first string
// 	last  string
// }

func main() {
	fmt.Println("Hello, world.")
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	totalSnow := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()
		ints := []int{}
		for index, char := range line {
			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				ints = append(ints, digit)
			} else if char == 'o' && index+2 <= len(line)-1 {
				if line[index+1] == 'n' && line[index+2] == 'e' {
					ints = append(ints, 1)
				} else if line[index+1] == 'n' && line[index+2] == 'e' && line[index+3] == 'r' {
					ints = append(ints, 0)
				}
			} else if char == 't' && index+2 <= len(line)-1 {
				if line[index+1] == 'w' && line[index+2] == 'o' {
					ints = append(ints, 2)
				} else if line[index+1] == 'h' && line[index+2] == 'r' && line[index+3] == 'e' && line[index+4] == 'e' {
					ints = append(ints, 3)
				}
			} else if char == 'f' && index+3 <= len(line)-1 {
				if line[index+1] == 'o' && line[index+2] == 'u' && line[index+3] == 'r' {
					ints = append(ints, 4)
				} else if line[index+1] == 'i' && line[index+2] == 'v' && line[index+3] == 'e' {
					ints = append(ints, 5)
				}
			} else if char == 's' && index+3 <= len(line)-1 {
				if line[index+1] == 'i' && line[index+2] == 'x' {
					ints = append(ints, 6)
				} else if line[index+1] == 'e' && line[index+2] == 'v' && line[index+3] == 'e' && line[index+4] == 'n' {
					ints = append(ints, 7)
				}
			} else if char == 'e' && index+4 <= len(line)-1 {
				if line[index+1] == 'i' && line[index+2] == 'g' && line[index+3] == 'h' && line[index+4] == 't' {
					ints = append(ints, 8)
				}
			} else if char == 'n' && index+3 <= len(line)-1 {
				if line[index+1] == 'i' && line[index+2] == 'n' && line[index+3] == 'e' {
					ints = append(ints, 9)
				}
			}

		}

		tens := ints[0] * 10
		units := ints[len(ints)-1]
		fmt.Println(tens + units)
		totalSnow += tens + units
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalSnow)
	file.Close()
}
