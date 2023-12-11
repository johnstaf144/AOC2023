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

func compareScoresJ(i []int, j []int) bool {
	for k := range i {
		if i[k] != j[k] {
			return i[k] > j[k]
		}
	}
	return false
}

func jokerScore(hand string) []int {
	cards := "23456789TQKA"

	if strings.Contains(hand, "J") {

		best_score_joker := [][]int{}
		for _, c := range cards {
			best_score_joker = append(best_score_joker, jokerScore(strings.Replace(hand, "J", string(c), 1)))
		}
		sort.Slice(best_score_joker, func(i, j int) bool {
			return compareScoresJ(best_score_joker[i], best_score_joker[j])
		})
		return best_score_joker[0]
	}
	return score(hand)
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
	score := jokerScore(hand)
	for _, i := range hand {
		tiebreak_score = append(tiebreak_score, strings.Index("J23456789TQKA", string(i)))
	}
	return append(score, tiebreak_score...)
}
