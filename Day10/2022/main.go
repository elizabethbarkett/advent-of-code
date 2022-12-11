package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var values []int
var crt = 0
var cycle = 0
var x = 1

var crtRow []string

func incrementCycle() {
	cycle += 1
	if crt == x-1 || crt == x || crt == x+1 {
		crtRow = append(crtRow, "#")
	} else {
		crtRow = append(crtRow, ".")
	}
	if crt == 39 {
		crt = 0
	} else {
		crt += 1
	}
}

func cycleCheck() {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		values = append(values, cycle*x)
	}
}

func calculate(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		i := scanner.Text()
		d := strings.Split(i, " ")
		instruction := d[0]

		if instruction == "noop" {
			incrementCycle()
			continue
		}

		if instruction == "addx" {
			j := 1
			v, _ := strconv.Atoi(d[1])
			for j < 2 {
				incrementCycle()
				cycleCheck()
				j += 1
			}
			if j == 2 {
				incrementCycle()
				cycleCheck()
				x = x + v
				continue
			}
		}
	}
	sum := 0
	for _, v := range values {
		sum += v
	}

	for m, n := range crtRow {
		if (m)%40 == 0 {
			fmt.Println()
		}
		fmt.Print(n)
	}
}

func main() {
	calculate("input.txt")
}
