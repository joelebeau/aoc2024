package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			s := strings.Replace(match, "mul", "", -1)
			s = strings.Replace(s, "(", "", -1)
			s = strings.Replace(s, ")", "", -1)
			ns1 := strings.Split(s, ",")[0]
			ns2 := strings.Split(s, ",")[1]
			n1, e1 := strconv.Atoi(ns1)
			n2, e2 := strconv.Atoi(ns2)
			fmt.Println(match, s)
			if e1 != nil || e2 != nil {
				panic("BOOM")
			}

			total += n1 * n2
		}
	}

	fmt.Println("result, ", total)

	// fmt.Println("RESULT: ", total)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}
