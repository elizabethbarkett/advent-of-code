package utils

import (
	"strconv"
)

func StringtoInt(number string) int {
	i, _ := strconv.Atoi(number)
	return i
}

func IterateAndConvert(s [][]string) map[int]int {
	var d = make(map[int]int)
	for i, _ := range s {
		value := s[i][0]
		g := StringtoInt(value)
		d[i]=g
	}
	return d
}