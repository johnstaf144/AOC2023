package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	var result float64 = 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[0-9]+`)
	for scanner.Scan() {

		card := scanner.Text()

		card_split := strings.Split(card, ": ")
		numbers := strings.Split(card_split[1], "|")
		winning_nums := re.FindAllString(numbers[0], -1)
		my_nums := re.FindAllString(numbers[1], -1)

		var win_matches float64 = 0
		for _, num := range winning_nums {
			if slices.Contains(my_nums, num) {
				win_matches++
			}
		}
		if win_matches != 0 {
			result += math.Pow(2, (win_matches - 1))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
