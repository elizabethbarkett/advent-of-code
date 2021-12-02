package main

import (
	"advent-of-code/Day2/2021/utils"
	"fmt"
)
//forward adds to position
// up decreases depth
// down increases depth
// multiply position * depth to get final answer

func main () {
	m := utils.ReadCsvFile("input_day2.csv")
	newslice := utils.IterateAndParse(m)
	position := newslice["forward"]
	depth := newslice["down"] - newslice["up"]
	answer := depth * position
	fmt.Println(answer)
}
