package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Card struct {
	my_num_wins int
	card_count  int
}

func main() {
	result := 0

	cards := []*Card{}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[0-9]+`)
	for scanner.Scan() {
		my_num_wins := 0

		card := scanner.Text()

		card_split := strings.Split(card, ": ")
		numbers := strings.Split(card_split[1], "|")
		winning_nums := re.FindAllString(numbers[0], -1)
		my_nums := re.FindAllString(numbers[1], -1)

		for _, num := range winning_nums {
			if slices.Contains(my_nums, num) {
				my_num_wins++
			}
		}

		cards = append(
			cards, &Card{
				my_num_wins: my_num_wins,
				card_count:  1,
			},
		)
	}

	for i, card := range cards {
		result += card.card_count
		for j := 1; j <= card.my_num_wins && i+j < len(cards); j++ {
			cards[i+j].card_count += card.card_count
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
