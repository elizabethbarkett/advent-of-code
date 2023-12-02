package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	for scanner.Scan() {
		num, err := partOne(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		sumPart1 += num
	}

	fmt.Println(sumPart1)
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

// func partTwo(text string) (int, error) {
// 	var current string
// 	var key map[string]int
// 	for _, c := range text {
// 		n, err := strconv.Atoi(string(c))

// 		if err != nil {
// 			letter := string(c)
// 			// start checking for a string digit
// 			if current == "" {
// 				m, ok := digits[letter]
// 				if !ok && current == "" {
// 					continue
// 				}
// 				key = m
// 			}
// 			// check to see if current exists in target map
// 			current = current + letter
// 			v, ok := key[current]
// 			if ok {
// 				return v, nil
// 			}
// 			for k, v := range key {

// 			}
// 		}
// 	}
// 	return 0, nil
// }

// var digits = map[string]map[string]int{
// 	"e": {
// 		"eight": 8,
// 	},
// 	"f": {
// 		"four": 4,
// 		"five": 5,
// 	},
// 	"n": {
// 		"nine": 9,
// 	},
// 	"o": {
// 		"one": 1,
// 	},
// 	"s": {
// 		"six":   6,
// 		"seven": 7,
// 	},
// 	"t": {
// 		"two":   2,
// 		"three": 3,
// 	},
// }
