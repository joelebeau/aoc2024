package main

// This is a little different than my original p6-1 solution.
// I forgot to copy things over and do part 2 separately, so I
// started my refactor for part 2, but just made sure it worked
// for part 1 before moving on to the part 2 solution. Oops.

import (
	"bufio"
	"fmt"
	"os"
)

// Simplified "pseudo" edges
type Node struct {
	Right      *Node
	Left       *Node
	Up         *Node
	Down       *Node
	Obstructed bool
	Visited    bool
	Coords     [2]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	tCount := run(lines)

	fmt.Println("total: ", tCount)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}

func run(lines []string) int {
	var startNode *Node
	endNode := &Node{
		Right:  nil,
		Left:   nil,
		Up:     nil,
		Down:   nil,
		Coords: [2]int{-1, -1},
	}

	nodes := [][]*Node{}
	for y := 0; y < len(lines); y++ {
		line := lines[y]
		nodes = append(nodes, []*Node{})
		for x, r := range line {
			obstructed := r == '#'
			visited := r == '^'

			n := Node{
				Right:      nil,
				Left:       nil,
				Up:         nil,
				Down:       nil,
				Obstructed: obstructed,
				Visited:    visited,
				Coords:     [2]int{y, x},
			}

			if visited {
				startNode = &n
			}

			nodes[y] = append(nodes[y], &n)
			if visited {
				startNode = &n
			}
		}
	}
	fmt.Println("Done building nodes")

	// join nodes
	fmt.Println("Linking nodes")
	for yIdx, row := range nodes {
		for xIdx, node := range row {
			if xIdx == len(row)-1 {
				node.Right = endNode
			} else {
				node.Right = nodes[yIdx][xIdx+1]
			}

			if xIdx == 0 {
				node.Left = endNode
			} else {
				node.Left = nodes[yIdx][xIdx-1]
			}

			if yIdx == len(nodes)-1 {
				node.Down = endNode
			} else {
				node.Down = nodes[yIdx+1][xIdx]
			}

			if yIdx == 0 {
				node.Up = endNode
			} else {
				node.Up = nodes[yIdx-1][xIdx]
			}
			fmt.Println(node)
		}
	}
	fmt.Println("Done linking nodes")

	tCount := traverse(startNode, endNode)

	return tCount
}

func traverse(startNode *Node, endNode *Node) int {
	dir := "north"
	traversed := 1
	node := startNode
	for node != endNode {
		var next *Node = nil
		for next == nil || next.Obstructed == true {
			if next != nil && next.Obstructed == true {
				dir = nextDir(dir)
			}
			switch dir {
			case "north":
				next = node.Up
			case "east":
				next = node.Right
			case "south":
				next = node.Down
			case "west":
				next = node.Left
			}
		}

		node = next
		if next == endNode {
			break
		} else {
			if !node.Visited {
				node.Visited = true
				traversed++
			}
		}
	}

	return traversed
}

func nextDir(dir string) string {
	switch dir {
	case "north":
		return "east"
	case "east":
		return "south"
	case "south":
		return "west"
	case "west":
		return "north"
	}

	return "north"
}
