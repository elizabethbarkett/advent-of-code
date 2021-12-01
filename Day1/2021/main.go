package main

import (
	"Day1/utils"
	"fmt"
)

func main() {
	s := utils.ReadCsvFile("input_day1.csv")
	m := utils.IterateAndConvert(s)
	partOne := 0
	partTwo := 0
	for i, v := range(m){
		if m[i+1] > v {
			partOne+=1
		}
	}
	fmt.Println(partOne)
	for i, v := range(m){
		w1 := v + m[i+1] + m[i+2]
		w2 := m[i+1] + m[i+2] + m[i+3]
		if w2 > w1 {
			partTwo+=1
		}
	}
	fmt.Println(partTwo)
}
