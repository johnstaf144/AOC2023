package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := 0

	actual_bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		valid := true

		line_split := strings.Split(line, ":")

		re := regexp.MustCompile("[0-9]+")

		game_number, _ := strconv.Atoi(re.FindAllString(line_split[0], -1)[0])

		ball_sets := strings.Split(line_split[1], ";")

		for _, set := range ball_sets {
			balls_in_set := strings.Split(set, ",")
			for _, ball_type := range balls_in_set {
				ball_info := strings.Split(ball_type, " ")
				colour := ball_info[2]
				number_of_balls, _ := strconv.Atoi(ball_info[1])
				if number_of_balls > actual_bag[colour] {
					valid = false
				}
			}
		}

		if valid {
			result += game_number
		}
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
