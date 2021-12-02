package utils

import (
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

func IterateAndParse(s [][]string) map[string]int {
	var d = make(map[string]int)
	aim := 0
	for i, _ := range s {
		value := s[i][0]
		ns := strings.Split(value, " ")
		key := ns[0]
		g := StringtoInt(ns[1])
		//d[key]=d[key] + g
		switch key {
		case "forward": d["down"] = d["down"]+(aim*g); d["forward"] = d["forward"]+g
		case "down": aim = aim+g
		case "up": aim = aim-g
		}
		d["aim"] = aim
	}
	return d
}