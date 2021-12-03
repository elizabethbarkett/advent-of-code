package main

import (
	"advent-of-code/utils"
	"fmt"
	"os"
)

func main() {
	dir, _:= os.Getwd()
	m := utils.ReadCsvFile(dir+"/input_day3.csv")
	l := len(m[0][0])
	gamma := ""
	epsilon := ""
	for i:=0; i < l; i++ {
		bit := utils.MostCommonBit(m, i, i+1)
		var bitE string
		if bit == "0" {
			bitE = "1"
		} else {
			bitE = "0"
		}
		gamma = gamma+bit
		epsilon = epsilon+bitE
	}
	numG, numE:= utils.StrBinaryToDecimal(gamma), utils.StrBinaryToDecimal(epsilon)
	answer := numG*numE
	fmt.Println("Answer: ", answer)
}