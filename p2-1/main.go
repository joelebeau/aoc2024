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
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, " ")
		nums := []int{}

		for _, num := range numsStr {
			x, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums = append(nums, x)
		}

		if isSafe(nums) {
			total++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}

	fmt.Println("RESULT: ", total)
}

func isSafe(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	if nums[1] > nums[0] {
		return isSafeIncreasing(nums)
	} else if nums[1] < nums[0] {
		return isSafeDecreasing(nums)
	}

	return false
}

func isSafeIncreasing(nums []int) bool {
	for idx, x := range nums {
		if idx == 0 {
			continue
		}

		if nums[idx-1] >= x {
			// Not all increasing
			return false
		}
		if x-nums[idx-1] > 3 {
			// Jump too big
			return false
		}
	}

	return true
}

func isSafeDecreasing(nums []int) bool {
	for idx, x := range nums {
		if idx == 0 {
			continue
		}

		if nums[idx-1] <= x {
			// Not all increasing
			return false
		}
		if nums[idx-1]-x > 3 {
			// Jump too big
			return false
		}
	}

	return true
}
