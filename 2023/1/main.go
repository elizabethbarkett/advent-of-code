package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sumPart1 int
	var sumPart2 int

	for scanner.Scan() {
		text := scanner.Text()
		num, err := partOne(text)
		if err != nil {
			fmt.Println(err)
			return
		}
		sumPart1 += num

		num, err = partTwo(text)
		if err != nil {
			fmt.Println(err)
			return
		}
		sumPart2 += num
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}

func partOne(text string) (int, error) {
	var addends []int
	for _, c := range text {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		addends = append(addends, n)
	}
	if len(addends) == 1 {
		addends = append(addends, addends[0])
	}
	fd := strconv.Itoa(addends[0])
	ld := strconv.Itoa(addends[len(addends)-1])
	return strconv.Atoi(fd + ld)
}

func partTwo(text string) (int, error) {

	var addends []int
	for i, c := range text {
		// if a number
		n, err := strconv.Atoi(string(c))
		if err != nil {
			digit := isNumber(text[i:])
			if digit != 0 {
				addends = append(addends, digit)
			}
		} else {
			addends = append(addends, n)
		}
	}
	if len(addends) == 1 {
		addends = append(addends, addends[0])
	}
	fd := strconv.Itoa(addends[0])
	ld := strconv.Itoa(addends[len(addends)-1])
	return strconv.Atoi(fd + ld)
}

func isNumber(text string) int {
	var current string
	var key map[string]int
	for _, k := range text {
		_, err := strconv.Atoi(string(k))
		if err != nil {
			if current == "" {
				m, ok := digits[string(k)]
				if !ok {
					current = ""
					continue
				}
				key = m
			}
			current = current + string(k)
			// does it exist?
			for k, v := range key {
				if strings.HasPrefix(current, k) {
					return v
				}
			}
		} else {
			return 0
		}
	}
	return 0
}

var digits = map[string]map[string]int{
	"e": {
		"eight": 8,
	},
	"f": {
		"four": 4,
		"five": 5,
	},
	"n": {
		"nine": 9,
	},
	"o": {
		"one": 1,
	},
	"s": {
		"six":   6,
		"seven": 7,
	},
	"t": {
		"two":   2,
		"three": 3,
	},
}
