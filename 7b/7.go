package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	value int
	cards string
}
type HandRank struct {
	Name string
	Rank int
}

func findMaxInMap(m map[string]int) string {
	var maxValue int
	var maxKey string
	duplicateKeys := []string{}
	for key, value := range m {
		if value > maxValue && key != "J" {
			maxValue = value
			maxKey = key
			duplicateKeys = []string{}
		} else if value == maxValue {
			duplicateKeys = append(duplicateKeys, key)
		}
	}
	if len(duplicateKeys) > 1  {
		sort.Slice(duplicateKeys, func(i, j int) bool {
			cardsRanking := map[string]int{"A": 12, "K": 11, "Q": 10, "T": 8, "9": 7, "8": 6, "7": 5, "6": 4, "5": 3, "4": 2, "3": 1, "2": 0, "J": -1}
			return cardsRanking[duplicateKeys[i]] > cardsRanking[duplicateKeys[j]]
		})
		return duplicateKeys[0]
	}
	return maxKey
}

func determineHandType(hand Hand) HandRank {
	cardCounts := make(map[string]int)
	same5 := 0
	same4 := 0
	same3 := 0
	same2 := 0
	same1 := 0

	for _, v := range hand.cards {
		cardCounts[string(v)] += 1
	}

	if cardCounts["J"] != 0 {
		maxKey := findMaxInMap(cardCounts)
		cardCounts[maxKey] += cardCounts["J"]
		if maxKey != "J" {
			cardCounts["J"] = 0
		}
	}

	for _, v := range cardCounts {
		if v == 5 {
			same5++
		}
		if v == 4 {
			same4++
		}
		if v == 3 {
			same3++
		}
		if v == 2 {
			same2++
		}
		if v == 1 {
			same1++
		}
	}

	if same5 == 1 {
		return HandRank{Name: "five of a kind", Rank: 0}
	} else if same4 == 1 {
		return HandRank{Name: "four of a kind", Rank: 1}
	} else if same3 == 1 && same2 == 1 {
		return HandRank{Name: "full house", Rank: 2}
	} else if same3 == 1 && same1 == 2 {
		return HandRank{Name: "three of a kind", Rank: 3}
	} else if same2 == 2 {
		return HandRank{Name: "two pair", Rank: 4}
	} else if same1 == 3 && same2 == 1 {
		return HandRank{Name: "one pair", Rank: 5}
	} else {
		return HandRank{Name: "high card", Rank: 6}
	}

}
func (h Hand) String() string {
	return fmt.Sprintf("Hand{Cards: %s, Value: %d}", h.cards, h.value)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	handBuckets := make([][]Hand, 7)
	totalHands := 0
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		value, _ := strconv.Atoi(line[1])
		hand := Hand{value: value, cards: line[0]}
		handRank := determineHandType(hand)
		handBuckets[handRank.Rank] = append(handBuckets[handRank.Rank], hand)
		totalHands++
	}

	//sorts in descending order hence the !true and !false
	for bucketIndex := range handBuckets {
		sort.Slice(handBuckets[bucketIndex], func(i, j int) bool {
			cardsRanking := map[string]int{"A": 12, "K": 11, "Q": 10, "T": 8, "9": 7, "8": 6, "7": 5, "6": 4, "5": 3, "4": 2, "3": 1, "2": 0, "J": -1}
			for k := 0; k < len(handBuckets[bucketIndex][i].cards); k++ {
				if cardsRanking[string(handBuckets[bucketIndex][i].cards[k])] < cardsRanking[string(handBuckets[bucketIndex][j].cards[k])] {
					return !true
				} else if cardsRanking[string(handBuckets[bucketIndex][i].cards[k])] > cardsRanking[string(handBuckets[bucketIndex][j].cards[k])] {
					return !false
				}
			}
			return false
		})
	}
	for bucketIndex := range handBuckets {
		for _, hand := range handBuckets[bucketIndex] {
			fmt.Println(hand, determineHandType(hand), totalHands)
			total += hand.value * totalHands

			totalHands--
		}
	}
	println(total)

}
