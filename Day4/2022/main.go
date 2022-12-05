package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isMaxRange(eval []int, compare []int) bool {
	//returns true for equal pairs as well.
	if (eval[1]-eval[0]) > (compare[1]-compare[0]) || eval[0] == compare[0] && eval[1] == compare[1] {
		return true
	}
	return false
}

func isOverlapping(eval []int, compare []int) bool {
	if compare[1] < eval[0] || eval[1] < compare[0] {
		return false
	} else if compare[0] <= eval[0] && eval[0] <= compare[1] || eval[0] <= compare[0] && compare[0] <= eval[1] || eval[0] <= compare[1] && eval[1] <= compare[1] || eval[0] <= compare[0] && eval[1] <= compare[1] && eval[1] >= compare[0] {
		return true
	}
	return false
}

func getIntSlice(sl []string) []int {
	var res []int
	for _, v := range sl {
		num, _ := strconv.Atoi(v)
		res = append(res, num)
	}
	return res
}

func calculateAssignmentPairsPart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		i := scanner.Text()
		//format the strings
		pairs := strings.Split(i, ",")
		pairOne := strings.Split(pairs[0], "-")
		pairTwo := strings.Split(pairs[1], "-")
		pairOneInt := getIntSlice(pairOne)
		pairTwoInt := getIntSlice(pairTwo)

		//get max and its compliment
		//check of other pair is equal/more than compliment but less/equal to max

		if isMaxRange(pairOneInt, pairTwoInt) {
			if pairOneInt[0] <= pairTwoInt[0] && pairOneInt[1] >= pairTwoInt[1] {
				counter += 1
			}
		} else if isMaxRange(pairTwoInt, pairOneInt) {
			if pairTwoInt[0] <= pairOneInt[0] && pairTwoInt[1] >= pairOneInt[1] {
				counter += 1
			}
		}
	}
	fmt.Println(counter)
}

func calculatePart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		i := scanner.Text()
		pairs := strings.Split(i, ",")
		pairOne := strings.Split(pairs[0], "-")
		pairTwo := strings.Split(pairs[1], "-")

		pairOneInt := getIntSlice(pairOne)
		pairTwoInt := getIntSlice(pairTwo)

		if isOverlapping(pairOneInt, pairTwoInt) {
			counter += 1
		}
	}
	fmt.Println(counter)
}

func main() {
	calculateAssignmentPairsPart1("input.txt")
	calculatePart2("input.txt")
}
