package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadTextToSums(filename string) []int {
	//text lines are added and sum is added to slice, slice items separated by newline in txt file
	f, err := os.Open(filename)
	var l []int
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		i := scanner.Text()
		num, e := strconv.Atoi(i)
		if e != nil {
			l = append(l, sum)
			sum = 0
		} else {
			sum = sum + num
		}
	}
	return l
}
