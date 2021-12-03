package utils

import (
	"fmt"
	"strconv"
	"strings"
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

func SliceToMap(s [][]string) map[string]string {
	var d = make(map[string]string)
	for i, _ := range s {
		value := s[i][0]
		d[value]=value
	}
	fmt.Println(d)
	return d
}


func IterateAndParse(s [][]string) map[string]int {
	var d = make(map[string]int)
	aim := 0
	for i, _ := range s {
		value := s[i][0]
		ns := strings.Split(value, " ")
		key := ns[0]
		g := StringtoInt(ns[1])
		switch key {
		case "forward": d["down"] = d["down"]+(aim*g); d["forward"] = d["forward"]+g
		case "down": aim = aim+g
		case "up": aim = aim-g
		}
		d["aim"] = aim
	}
	return d
}

func MostCommonBit(s [][]string, x int, y int) string {
	count0 := 0
	count1 := 0
	for i, _ := range(s) {
		value := s[i][0]
		firstChar := value[x:y]
		switch firstChar{
		case "0": count0+=1
		case "1": count1+=1
		}
	}
	if count1 >= count0 { return "1" } else {return "0"}
}

func LeastCommonBitPt2 (s [][]string, x int, y int) string {
	count0 := 0
	count1 := 0
	for i, _ := range (s) {
		value := s[i][0]
		firstChar := value[x:y]
		switch firstChar {
		case "0":
			count0 += 1
		case "1":
			count1 += 1
		}
	}
	if count0 <= count1 {
		return "0"
	} else {
		return "1"
	}
}

func IsMatch(s string, x int, y int, bit string) bool {
	value := s[x:y]
	if value == bit {
		return true
	}
	return false
}

func StrBinaryToDecimal(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 64)
	return i
}

func IsolateRatingFromMostCommonBit(bit string, s [][]string, x int, y int) [][]string {
	i := 0
	for i <= len(s)-1 {
		value := s[i][0]
		isMatch := IsMatch(value, x, y, bit)
		if isMatch { i+= 1}
		if !isMatch{
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}