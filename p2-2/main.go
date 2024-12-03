package main

// I made some dumb errors in this early on that took me a while to catch, so this turned into
// a little bit of a mess as I tried to figure out wtf I did wrong.
import (
	"bufio"
	"fmt"
	"math"
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
		} else {
			safeWithoutOne := false
			for idx := range nums {
				if isSafe(removeIdx(nums, idx)) {
					safeWithoutOne = true
				}
			}
			if safeWithoutOne {
				total++
			}
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

	if nums[1] == nums[0] {
		return false
	}

	if nums[1] > nums[0] {
		return isSafeIncreasing(nums)
	} else {
		return isSafeDecreasing(nums)
	}
}

func isSafeIncreasing(nums []int) bool {
	for idx, x := range nums {
		if idx == 0 {
			continue
		}

		if x <= nums[idx-1] {
			return false
		}
		if math.Abs(float64(x-nums[idx-1])) > 3 {
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

		if x >= nums[idx-1] {
			return false
		}
		if math.Abs(float64(x-nums[idx-1])) > 3 {
			return false
		}
	}

	return true
}

func removeIdx(s []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:i]...)
	if i+1 >= len(s) {
		return ret
	}
	return append(ret, s[i+1:]...)
}
