package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x, y, steps int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")

	paths := strings.Split(string(input), "\n")
	pathA := strings.Split(paths[0], ",")
	pathB := strings.Split(paths[1], ",")

	// fmt.Println(pathA)
	// fmt.Println("---")
	// fmt.Println(pathB)
	// fmt.Println("---")
	pathAxy := buildPath(pathA)
	pathBxy := buildPath(pathB)

	// fmt.Println(pathAxy)
	// fmt.Println("---")
	// fmt.Println(pathBxy)

	// finding crossings and steps count
	// var crossings []pos
	var stepTotals []int
	for _, posA := range pathAxy {
		for _, posB := range pathBxy {
			if posA.x == posB.x && posA.y == posB.y {
				// crossings = append(crossings, posA)
				// crossings = append(crossings, posB)
				stepTotals = append(stepTotals, posA.steps+posB.steps)
			}
		}
	}

	// fmt.Println(crossings)
	fmt.Println(stepTotals)

	// finding most closes to start point by Manhattan distance. Test input 610
	lowest := 1<<63 - 1
	for _, stepTotal := range stepTotals {
		if stepTotal < lowest {
			lowest = stepTotal
		}
	}

	fmt.Println(lowest)
}

// create all points for the path
func buildPath(input []string) []pos {
	var path []pos
	var steps int
	x, y := 0, 0
	for _, moveCmd := range input {
		offset, _ := strconv.Atoi(moveCmd[1:])
		switch string(moveCmd[0]) {
		case "R":
			for j := 0; j < offset; j++ {
				x++
				steps++
				path = append(path, pos{x, y, steps})
			}
		case "L":
			for j := 0; j < offset; j++ {
				x--
				steps++
				path = append(path, pos{x, y, steps})
			}
		case "U":
			for j := 0; j < offset; j++ {
				y++
				steps++
				path = append(path, pos{x, y, steps})
			}
		case "D":
			for j := 0; j < offset; j++ {
				y--
				steps++
				path = append(path, pos{x, y, steps})
			}
		}
	}
	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
