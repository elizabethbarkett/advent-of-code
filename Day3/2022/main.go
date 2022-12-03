package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getValue(c rune) int {
	if c > 96 {
		return int(c) - 96
	}
	return int(c) - 38
}

func stringToSlice(s string) []string {
	var r []string
	for _, v := range s {
		r = append(r, string(v))
	}
	return r
}

func Contains(j []rune, searchValue rune) bool {
	for _, value := range j {
		if searchValue != 44 && value == searchValue {
			return true
		}
	}
	return false
}

func calculateItemsPart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		l := len(i)
		s1 := stringToSlice(i[0:(l / 2)])
		s2 := stringToSlice(i[(l / 2):l])
		sort.Strings(s1)
		sort.Strings(s2)
		rs1 := []rune(strings.Join(s1[:], ","))
		rs2 := []rune(strings.Join(s2[:], ","))

		for i := 0; i <= (len(rs1)); i++ {
			currentValue := rs1[i]
			if Contains(rs2, rs1[i]) {
				currentValuePoints := getValue(currentValue)
				points = points + currentValuePoints
				break
			}
		}
	}
	fmt.Println(points)
}

func calculateItemsPart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	group := make(map[int][]rune)
	points := 0
	indexLine := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//groups of 3 lines
		i := scanner.Text()
		s1 := stringToSlice(i)
		sort.Strings(s1)
		rs1 := []rune(strings.Join(s1[:], ","))
		group[indexLine] = rs1
		indexLine = indexLine + 1
	}

	for i := 0; i < len(group)-2; i += 3 {
		slice1 := group[i]
		slice2 := group[i+1]
		slice3 := group[i+2]

		groupMap := map[int][]rune{1: slice1, 2: slice2, 3: slice3}
		keys := make([]int, 0, len(groupMap))

		for k, _ := range groupMap {
			keys = append(keys, k)
		}

		var sortedList [][]rune

		for _, k := range keys {
			sortedList = append(sortedList, groupMap[k])
		}

		for _, value := range sortedList[0] {
			if Contains(sortedList[1], value) && Contains(sortedList[2], value) {
				points = points + getValue(value)
				break
			}
		}
	}
	fmt.Println(points)
}

func main() {
	calculateItemsPart1("input.txt")
	calculateItemsPart2("input.txt")
}
