package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		power := 1
		min_balls := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		line := scanner.Text()

		line_split := strings.Split(line, ":")

		ball_sets := strings.Split(line_split[1], ";")

		for _, set := range ball_sets {
			balls_in_set := strings.Split(set, ",")
			for _, ball_type := range balls_in_set {
				ball_info := strings.Split(ball_type, " ")
				colour := ball_info[2]
				number_of_balls, _ := strconv.Atoi(ball_info[1])

				if number_of_balls > min_balls[colour] {
					min_balls[colour] = number_of_balls
				}thisistestDaria
			}
		}

		for _, val := range min_balls {
			power = power * val
		}

		result += power
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
