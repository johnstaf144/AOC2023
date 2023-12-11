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

	seeds := []int{}
	seed_maps := [][][]int{}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	map_index := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			if re := regexp.MustCompile(`seeds`); re.MatchString(line) {
				for _, seed := range strings.Split(strings.Split(line, "seeds: ")[1], " ") {
					seed_num, _ := strconv.Atoi(seed)
					seeds = append(seeds, seed_num)
				}
			} else if re := regexp.MustCompile(`[0-9]+`); !re.MatchString(line) {
				seed_maps = append(seed_maps, [][]int{})
				map_index++
			} else {
				map_nums := []int{}
				for _, map_num_str := range strings.Split(line, " ") {
					map_num_int, _ := strconv.Atoi(map_num_str)
					map_nums = append(map_nums, map_num_int)
				}
				seed_maps[map_index-1] = append(seed_maps[map_index-1], map_nums)
			}
		}
	}

	for _, seed := range seeds {
		for _, i := range seed_maps {
			for _, j := range i {
				if j[1] <= seed && seed <= j[1]+j[2]-1 {
					seed += j[0] - j[1]
					break
				}
			}
		}
		if result == 0 || seed < result {
			result = seed
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
