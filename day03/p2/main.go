package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Num struct {
	int_num int
	i       int
	start   int
	j       int
}

func main() {
	result := 0
	board := []string{}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	BOARD_LEN := len(board)
	LINE_LEN := len(board[0])

	nums := []Num{}

	for i := 0; i < BOARD_LEN; i++ {
		j := 0
		for j < LINE_LEN {
			if _, err := strconv.Atoi(string(board[i][j])); err == nil {
				start := j
				num := ""
				for j < LINE_LEN {
					if _, err := strconv.Atoi(string(board[i][j])); err != nil {
						break
					}
					num += string(board[i][j])
					j += 1
				}
				j -= 1
				int_num, _ := strconv.Atoi(num)
				nums = append(
					nums,
					Num{
						int_num: int_num,
						i:       i,
						start:   start,
						j:       j,
					},
				)
			}
			j += 1
		}
	}

	gear_map := map[string][]int{}

	for _, num := range nums {
		for i := num.i - 1; i <= num.i+1; i++ {
			if i >= 0 && i < BOARD_LEN {
				for j := num.start - 1; j <= num.j+1; j++ {
					if j >= 0 && j < LINE_LEN {
						if string(board[i][j]) == "*" {
							gear_index := fmt.Sprintf("%d,%d", i, j)
							if val, isKeyExists := gear_map[gear_index]; !isKeyExists {
								gear_map[gear_index] = []int{num.int_num}
							} else {
								gear_map[gear_index] = append(val, num.int_num)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(gear_map)

	for _, val := range gear_map {
		if len(val) == 2 {
			result += val[0] * val[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
