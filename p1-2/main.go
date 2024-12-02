package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, "   ")
		lInt, err := strconv.Atoi(segments[0])
		if err != nil {
			panic(err)
		}
		rInt, err := strconv.Atoi(segments[1])
		if err != nil {
			panic(err)
		}
		left = append(left, lInt)
		right = append(right, rInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}

	total := 0
	rMap := make(map[int]int)
	for _, x := range right {
		rMap[x]++
	}
	for _, x := range left {
		total += x * rMap[x]
	}

	fmt.Println("RESULT: ", total)
}
