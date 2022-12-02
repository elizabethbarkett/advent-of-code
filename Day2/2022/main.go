package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getWinner(prompt rune) rune {
	if prompt != 67 {
		return prompt + 24
	} else {
		return prompt + 21
	}
}

func getLoser(prompt rune) rune {
	if prompt != 65 {
		return prompt + 22
	} else {
		return prompt + 25
	}
}

func getCompliment(prompt rune) rune {
	return prompt + 23
}

func getPoints(response rune) int {
	switch response {
	case 88:
		return 1
	case 89:
		return 2
	case 90:
		return 3
	}
	return 0
}

func getResultsPart1(filename string) int {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		prompt := rune(i[0])
		response := rune(i[2])
		w := getWinner(prompt)
		c := getCompliment(prompt)
		responsePoints := getPoints(response)

		if response == w {
			points = points + responsePoints + 6
		} else if response == c {
			points = points + responsePoints + 3
		} else {
			points = points + responsePoints + 0
		}

	}
	fmt.Println(points)
	return points
}

func getResultsPart2(filename string) int {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		prompt := rune(i[0])
		result := rune(i[2])
		w := getWinner(prompt)
		c := getCompliment(prompt)
		l := getLoser(prompt)
		if result == 88 {
			points = points + 0 + getPoints(l)
		} else if result == 89 {
			points = points + 3 + getPoints(c)
		} else {
			points = points + 6 + getPoints(w)
		}
	}
	fmt.Println(points)
	return points
}

func main() {
	getResultsPart1("input.txt")
	getResultsPart2("input.txt")
}
