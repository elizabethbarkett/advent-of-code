package main

import (
	"fmt"

	"./utils"
)

func main() {
	filename := "./input.txt"
	s := utils.ReadTextToSums(filename)
	//PART 1
	setMax := 0
	for _, value := range s {
		if value > setMax {
			setMax = value
		}
	}
	fmt.Print(setMax, "\n")

	//PART 2
	var topThree [3]int
	for _, value := range s {
		if value > topThree[0] {
			topThree[2] = topThree[1]
			topThree[1] = topThree[0]
			topThree[0] = value
		} else if topThree[0] > value && value > topThree[1] {
			topThree[2] = topThree[1]
			topThree[1] = value
		} else if topThree[1] > value && value > topThree[2] {
			topThree[2] = value
		}
	}
	fmt.Print(topThree[0] + topThree[1] + topThree[2])
}
