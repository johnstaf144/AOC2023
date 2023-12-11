package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var result float64 = 1

	race_map := map[string][]float64{
		"time":     {},
		"distance": {},
	}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re_nums := regexp.MustCompile(`[0-9]+`)

	re_time := regexp.MustCompile(`Time`)

	for scanner.Scan() {
		line := scanner.Text()

		fixed_num_str := ""

		if re_time.MatchString(line) {
			nums_strings := re_nums.FindAllString(line, -1)
			for _, num_str := range nums_strings {
				fixed_num_str += num_str
			}
			num, _ := strconv.Atoi(fixed_num_str)
			race_map["time"] = append(race_map["time"], float64(num))
		} else {
			nums_strings := re_nums.FindAllString(line, -1)
			for _, num_str := range nums_strings {
				fixed_num_str += num_str
			}
			num, _ := strconv.Atoi(fixed_num_str)
			race_map["distance"] = append(race_map["distance"], float64(num))
		}

	}

	var a, b, c float64

	for i, total_time := range race_map["time"] {
		a = -1
		b = total_time
		c = -race_map["distance"][i]

		d := b*b - 4*a*c
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)

		result = result * (math.Ceil(root2) - math.Floor(root1) - 1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(int(result))
}
