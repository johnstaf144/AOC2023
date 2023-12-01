package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result := 0

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile("[1-9]")
		digits_list := re.FindAllString(line, -1)

		first_digit := digits_list[0]

		var rev []string
		for _, n := range digits_list {
			rev = append([]string{n}, rev...)
		}

		last_digit := rev[0]

		concat_digits_int, _ := strconv.Atoi(first_digit + last_digit)

		result += concat_digits_int
	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
