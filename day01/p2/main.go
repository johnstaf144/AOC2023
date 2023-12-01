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

	number_dict := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for key, value := range number_dict {
			re := regexp.MustCompile(key)
			line = re.ReplaceAllString(line, value)
		}

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
