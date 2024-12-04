package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

const XMAS_LEN = 4
const MAX_X = 136
const MAX_Y = 136
const MIN_X = 3
const MIN_Y = 3

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
	if field[y][x] != 'X' {
		return 0
	}
	total := 0
	total += crawlH(field, x, y)
	total += crawlV(field, x, y)
	total += crawlD(field, x, y)

	return total
}

func crawlH(field [][]rune, x, y int) int {
	total := 0
	//forward
	if x <= MAX_X {
		if field[y][x+1] == 'M' &&
			field[y][x+2] == 'A' &&
			field[y][x+3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping HF", x, y)
	}

	//back
	if x >= MIN_X {
		if field[y][x-1] == 'M' &&
			field[y][x-2] == 'A' &&
			field[y][x-3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping HB", x, y)
	}

	return total
}

func crawlV(field [][]rune, x, y int) int {
	total := 0
	//up
	if y >= MIN_Y {
		if field[y-1][x] == 'M' &&
			field[y-2][x] == 'A' &&
			field[y-3][x] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping VU", x, y)
	}

	//down
	if y <= MAX_Y {
		if field[y+1][x] == 'M' &&
			field[y+2][x] == 'A' &&
			field[y+3][x] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping VD", x, y)
	}

	return total
}

func crawlD(field [][]rune, x, y int) int {
	total := 0

	//up-right
	if x <= MAX_X && y >= MIN_Y {
		if field[y-1][x+1] == 'M' &&
			field[y-2][x+2] == 'A' &&
			field[y-3][x+3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping DUR", x, y)
	}

	//up-left
	if x >= MIN_X && y >= MIN_Y {
		if field[y-1][x-1] == 'M' &&
			field[y-2][x-2] == 'A' &&
			field[y-3][x-3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping DUL", x, y)
	}

	//down-right
	if x <= MAX_X && y <= MAX_Y {
		if field[y+1][x+1] == 'M' &&
			field[y+2][x+2] == 'A' &&
			field[y+3][x+3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping DDR", x, y)
	}

	//down-left
	if x >= MIN_X && y <= MAX_Y {
		if field[y+1][x-1] == 'M' &&
			field[y+2][x-2] == 'A' &&
			field[y+3][x-3] == 'S' {
			total++
		}
	} else {
		// fmt.Println("Skipping DDL", x, y)
	}

	return total
}
