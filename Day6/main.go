package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func evaluate(l []rune) bool {
	fmt.Println(l)
	for index, value := range l {
		//get first number and compare to last 3
		//get second number and compare to last 2
		//get third number and compare to fourth
		for _, v := range l[index+1:] {
			if value == v {
				return false
			}
			continue
		}
	}
	return true
}

func calculatePart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		i := scanner.Text()
		var eval []rune
		for index, r := range i {
			eval = append(eval, r)
			if len(eval) == 4 && evaluate(eval) {
				counter = index + 1 //because we are counting items not the index
				break
			} else if len(eval) == 4 && !evaluate(eval) {
				eval = eval[1:]
			} else {
				continue
			}
		}
	}
	fmt.Println(counter)
}

func calculatePart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		i := scanner.Text()
		var eval []rune
		for index, r := range i {
			eval = append(eval, r)
			if len(eval) == 14 && evaluate(eval) {
				counter = index + 1
				break
			} else if len(eval) == 14 && !evaluate(eval) {
				eval = eval[1:]
			} else {
				continue
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	calculatePart1("input.txt")
	calculatePart2("input.txt")
}
