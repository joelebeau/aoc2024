package main

import (
	"bufio"
	"fmt"
	"os"
)

const XMAS_LEN = 4

const MAX_X = 138
const MAX_Y = 138
const MIN_X = 1
const MIN_Y = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	field := make([][]rune, len(lines))
	y := 0
	for _, line := range lines {
		x := 0
		field[y] = make([]rune, len(line))
		for _, r := range line {
			field[y][x] = r
			x++
		}
		y++
	}
	total := 0
	for yIdx, line := range field {
		for xIdx := range line {
			total += crawl(field, xIdx, yIdx)
		}
	}
	fmt.Println("TOTAL: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}

func crawl(field [][]rune, x, y int) int {
	if field[y][x] != 'A' {
		return 0
	}
	total := 0
	total += crawlD(field, x, y)

	return total
}

func crawlD(field [][]rune, x, y int) int {
	found := map[rune]bool{
		'M': false,
		'S': false,
	}

	// PAIR 1
	//up-right
	if x <= MAX_X && y >= MIN_Y {
		found[field[y-1][x+1]] = true
	}

	//down-left
	if x >= MIN_X && y <= MAX_Y {
		found[field[y+1][x-1]] = true
	}

	dirOneFound := true
	for _, v := range found {
		if !v {
			dirOneFound = false
		}
	}
	if !dirOneFound {
		return 0
	}

	// PAIR 2
	found['M'] = false
	found['S'] = false
	//up-left
	if x >= MIN_X && y >= MIN_Y {
		found[field[y-1][x-1]] = true
	}

	//down-right
	if x <= MAX_X && y <= MAX_Y {
		found[field[y+1][x+1]] = true
	}
	dirTwoFound := true
	for _, v := range found {
		if !v {
			dirTwoFound = false
		}
	}
	if dirOneFound && dirTwoFound {
		return 1
	}

	return 0
}
