package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// These are technically numbers...but they aren't used numerically
type Rule struct {
	Before string
	After  string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := []Rule{}
	updatesList := [][]string{}
	mode := "rules"
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mode = "updates"
			continue
		}

		if mode == "rules" {
			rule := Rule{}
			nums := strings.Split(line, "|")
			rule.Before = strings.TrimSpace(nums[0])
			rule.After = strings.TrimSpace(nums[1])
			rules = append(rules, rule)
		} else if mode == "updates" {
			updates := strings.Split(line, ",")
			updatesList = append(updatesList, updates)
		}
	}
	invalidUpdates := [][]string{}
	for _, u := range updatesList {
		if !areValidUpdates(rules, u) {
			invalidUpdates = append(invalidUpdates, u)
		}
	}

	fixedUpdates := [][]string{}
	for _, u := range invalidUpdates {
		fixedUpdates = append(fixedUpdates, fixUpdates(rules, u))
	}

	total := 0
	for _, u := range fixedUpdates {
		idx := len(u) / 2
		x, err := strconv.Atoi(u[idx])
		if err != nil {
			panic(err)
		}
		total += x
	}

	fmt.Println("total: ", total)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}

func areValidUpdates(rules []Rule, updates []string) bool {
	valid := true
	for idx, u := range updates {
		rulesForCurrentVal := findRules(rules, u)
		for _, r := range rulesForCurrentVal {
			if slices.Contains(updates[:idx], r.After) {
				valid = false
			}
		}
	}

	return valid
}

// Returns rules for the given before value.
func findRules(rules []Rule, before string) []Rule {
	out := []Rule{}
	for _, rule := range rules {
		if rule.Before == before {
			out = append(out, rule)
		}
	}

	return out
}

func fixUpdates(rules []Rule, updates []string) []string {
	valid := false
	for !valid {
		func() {
			for idx, u := range updates {
				rulesForCurrentVal := findRules(rules, u)
				for _, r := range rulesForCurrentVal {
					if idx2 := slices.Index(updates[:idx], r.After); idx2 != -1 {
						tmp := updates[idx]
						updates[idx] = updates[idx2]
						updates[idx2] = tmp

						return
					}
				}
			}
			valid = true
		}()
	}

	return updates
}
