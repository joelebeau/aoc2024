package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

	if len(left) != len(right) {
		panic("UNEQUAL LEN? HOW?")
	}

	slices.Sort(left)
	slices.Sort(right)
	total := 0
	for idx := range left {
		l := float64(left[idx])
		r := float64(right[idx])
		total += int(math.Abs(l - r))
	}
	fmt.Println("RESULT: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}
