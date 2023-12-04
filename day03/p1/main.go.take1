package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type engine struct {
	is_part bool
}

func main() {
	result := 0
	board := []string{}

	e := engine{
		is_part: false,
	}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	engine_number_list := []int{}

	number_store := ""

	re := regexp.MustCompile(`[0-9]|\.`)

	for row_index, row := range board {
		for char_index := 0; char_index < len(row); char_index++ {
			char := string(row[char_index])
			if _, err := strconv.Atoi(char); err == nil {
				number_store += char

				if row_index == 0 && char_index == 0 { // top left corner
					e.check_front(board, row_index, char_index, re)
					e.check_below_right(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					if e.is_part {
						continue
					}
				} else if row_index == 0 && char_index == len(row)-1 { // top right corner
					e.check_back(board, row_index, char_index, re)
					e.check_below_right(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					if e.is_part {
						continue
					}
				} else if row_index == len(board)-1 && char_index == len(row)-1 { // bottom right corner (last)
					e.check_back(board, row_index, char_index, re)
					e.check_above_left(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					if e.is_part {
						engine_number, _ := strconv.Atoi(number_store)
						engine_number_list = append(engine_number_list, engine_number)
						continue
					}
				} else if row_index == len(board)-1 && char_index == 0 { // bottom left corner
					e.check_front(board, row_index, char_index, re)
					e.check_above_right(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					if e.is_part {
						continue
					}
				} else if row_index == 0 && char_index != 0 && char_index != len(row)-1 { // top of columns excl corners
					e.check_front(board, row_index, char_index, re)
					e.check_back(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					e.check_below_left(board, row_index, char_index, re)
					e.check_below_right(board, row_index, char_index, re)
					if e.is_part {
						continue
					}

				} else if row_index == len(board)-1 && char_index != 0 && char_index != len(row)-1 { // bottom of columns excl corners
					e.check_front(board, row_index, char_index, re)
					e.check_back(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					e.check_above_left(board, row_index, char_index, re)
					e.check_above_right(board, row_index, char_index, re)
					if e.is_part {
						continue
					}
				} else if row_index != 0 && row_index != len(board)-1 && char_index == 0 { // start of row excl corners
					e.check_front(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					e.check_above_right(board, row_index, char_index, re)
					e.check_below_right(board, row_index, char_index, re)
					if e.is_part {
						continue
					}

				} else if row_index != 0 && row_index != len(board)-1 && char_index == len(row)-1 { // end of row excl corners
					e.check_back(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					e.check_above_left(board, row_index, char_index, re)
					e.check_below_left(board, row_index, char_index, re)
					if e.is_part {
						continue
					}

				} else {
					e.check_front(board, row_index, char_index, re)
					e.check_back(board, row_index, char_index, re)
					e.check_above(board, row_index, char_index, re)
					e.check_below(board, row_index, char_index, re)
					e.check_above_left(board, row_index, char_index, re)
					e.check_above_right(board, row_index, char_index, re)
					e.check_below_left(board, row_index, char_index, re)
					e.check_below_right(board, row_index, char_index, re)
					if e.is_part {
						continue
					}
				}
			}

			if _, err := strconv.Atoi(char); err != nil || (row_index == len(board)-1 && char_index == len(row)-1) {
				if number_store != "" && e.is_part {
					engine_number, _ := strconv.Atoi(number_store)
					engine_number_list = append(engine_number_list, engine_number)
				}
				number_store = ""
				e.is_part = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, item := range engine_number_list {
		result += item
	}

	fmt.Println(result)
}

func (e *engine) check_front(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index][char_index+1])) {
		e.is_part = true
	}
}

func (e *engine) check_back(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index][char_index-1])) {
		e.is_part = true
	}
}

func (e *engine) check_above(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index-1][char_index])) {
		e.is_part = true
	}
}

func (e *engine) check_below(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index+1][char_index])) {
		e.is_part = true
	}
}

func (e *engine) check_above_left(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index-1][char_index-1])) {
		e.is_part = true
	}
}

func (e *engine) check_above_right(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index-1][char_index+1])) {
		e.is_part = true
	}
}

func (e *engine) check_below_left(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index+1][char_index-1])) {
		e.is_part = true
	}
}

func (e *engine) check_below_right(board []string, row_index int, char_index int, re *regexp.Regexp) {
	if !re.MatchString(string(board[row_index+1][char_index+1])) {
		e.is_part = true
	}
}
