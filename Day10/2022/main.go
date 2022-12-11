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
var cycle = 0
var x = 1

func cycleCheck() {
	//cycle += 1
	fmt.Println("Cycle: ", cycle)
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

		fmt.Println("X: ", x)

		i := scanner.Text()
		d := strings.Split(i, " ")
		instruction := d[0]

		if instruction == "noop" {
			cycle += 1
			continue
		}

		if instruction == "addx" {
			j := 1
			v, _ := strconv.Atoi(d[1])
			for j < 2 {
				cycle += 1
				cycleCheck()
				j += 1
			}
			if j == 2 {
				cycle += 1
				cycleCheck()
				x = x + v
				continue
			}
		}
	}
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	calculate("input.txt")
}
