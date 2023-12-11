package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ScoredHand struct {
	tiebreak []int
	hand     string
	bid      int
}

func main() {
	result := 0

	scored_hands := []ScoredHand{}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		hand_bid := strings.Split(line, " ")

		bid, _ := strconv.Atoi(hand_bid[1])

		scored_hands = append(scored_hands, ScoredHand{
			tiebreak: tiebreak(hand_bid[0]),
			hand:     hand_bid[0],
			bid:      bid,
		})

	}

	sort.Slice(scored_hands, func(i, j int) bool {
		return compareScores(scored_hands[i], scored_hands[j])
	})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	counter := 1

	for _, scored_hand := range scored_hands {
		result += (int(scored_hand.bid) * counter)
		counter++
	}

	fmt.Println(result)
}

func compareScores(i ScoredHand, j ScoredHand) bool {
	for k := range i.tiebreak {
		if i.tiebreak[k] != j.tiebreak[k] {
			return i.tiebreak[k] < j.tiebreak[k]
		}
	}
	return false
}

func score(hand string) []int {
	counts := map[string]int{}
	for _, card := range hand {
		_, keyExists := counts[string(card)]
		if !keyExists {
			counts[string(card)] = 0
		}
		counts[string(card)] += 1
	}
	values := make([]int, 0, len(counts))

	for _, value := range counts {
		values = append(values, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	return values
}

func tiebreak(hand string) []int {
	tiebreak_score := []int{}
	score := score(hand)
	for _, i := range hand {
		tiebreak_score = append(tiebreak_score, strings.Index("23456789TJQKA", string(i)))
	}
	return append(score, tiebreak_score...)
}
