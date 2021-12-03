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
	duplicate := make([][]string, len(m))
	for i := range m {
		duplicate[i] = make([]string, len(m[i]))
		copy(duplicate[i], m[i])
	}
	for i:=0; i < l; i++ {
		OxyBit := utils.MostCommonBit(m, i, i+1)
		CO2Bit := utils.LeastCommonBitPt2(duplicate, i, i+1)
		m = utils.IsolateRatingFromMostCommonBit(OxyBit, m, i, i+1)
		if len(duplicate) > 1 {
			duplicate = utils.IsolateRatingFromMostCommonBit(CO2Bit, duplicate, i, i+1)
		}
	}
	numC, numO := utils.StrBinaryToDecimal(m[0][0]), utils.StrBinaryToDecimal(duplicate[0][0])
	answer := numO*numC
	fmt.Println("Answer: ", answer)
}
