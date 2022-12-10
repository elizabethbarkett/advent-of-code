package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parsePuzzleMap(input []string) map[int][]string {
	m := make(map[int][]string)

	for index, value := range input {
		var n []string

		for i, v := range value {
			if (i+3)%4 == 0 && index != 8 {
				n = append(n, string(v))
			}
		}
		for i, v := range n {
			if i != 9 {
				m[i+1] = append(m[i+1], v)
			}
		}
	}
	return m
}

func parseMovement(input []string) [][]int {
	var set [][]int
	for _, j := range input {
		var l []int
		dir := strings.Split(j, " ")
		movement, _ := strconv.Atoi(dir[1])
		from, _ := strconv.Atoi(dir[3])
		to, _ := strconv.Atoi(dir[5])
		l = append(l, movement, from, to)
		set = append(set, l)
	}
	return set
}

func calculateMoves(set [][]int, puzzle map[int][]string) map[int][]string {

	for _, l := range set {
		num := l[0]
		from := l[1]
		to := l[2]
		i := 1
		for i <= num {
			el := puzzle[from][0]
			puzzle[from] = puzzle[from][1:]
			if len(puzzle[to]) == 0 {
				puzzle[to] = append(puzzle[to], el)
			}
			puzzle[to] = append(puzzle[to][:0+1], puzzle[to][0:]...)
			puzzle[to][0] = el
			i += 1
		}
	}
	return puzzle
}

func calculateMovesPart2(set [][]int, puzzle map[int][]string) map[int][]string {
	for _, l := range set {
		num := l[0]
		from := l[1]
		to := l[2]
		var el []string
		i := 0
		for i < num {
			el = append(el, puzzle[from][:num][i])
			i += 1
		}
		puzzle[from] = puzzle[from][num:]
		puzzle[to] = append(el, puzzle[to]...)
	}
	return puzzle
}

func cleanList(p map[int][]string) map[int][]string {
	for i, j := range p {
		for _, k := range j {
			if k == " " {
				j = j[1:]
			}
		}
		p[i] = j
	}
	return p
}

func calculate(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var puzzle []string
	var directions []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		puzzle = append(puzzle, i)
	}
	directions = puzzle[10:]
	puzzle = puzzle[:10]
	set := parseMovement(directions)
	fpuzzle := parsePuzzleMap(puzzle)
	fpuzzle = cleanList(fpuzzle)
	//fpuzzle = calculateMoves(set, fpuzzle)
	fpuzzle = calculateMovesPart2(set, fpuzzle)

	fmt.Println(fpuzzle)
}

func main() {
	calculate("input.txt")
}
