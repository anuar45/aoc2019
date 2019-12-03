package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
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

	// finding points of crossing
	var crossings []pos
	for _, posA := range pathAxy {
		for _, posB := range pathBxy {
			if posA.x == posB.x && posA.y == posB.y {
				crossings = append(crossings, posB)
			}
		}
	}

	fmt.Println(crossings)

	// finding most closes to start point by Manhattan distance. Test input 135
	lowest := 1<<63 - 1
	for _, cross := range crossings {
		lenght := abs(cross.x) + abs(cross.y)
		if lenght < lowest {
			lowest = lenght
		}
	}

	fmt.Println(lowest)
}

// create all points for the path
func buildPath(input []string) []pos {
	var path []pos
	x, y := 0, 0
	for _, moveCmd := range input {
		offset, _ := strconv.Atoi(moveCmd[1:])
		switch string(moveCmd[0]) {
		case "R":
			for j := 0; j < offset; j++ {
				x++
				path = append(path, pos{x, y})
			}
		case "L":
			for j := 0; j < offset; j++ {
				x--
				path = append(path, pos{x, y})
			}
		case "U":
			for j := 0; j < offset; j++ {
				y++
				path = append(path, pos{x, y})
			}
		case "D":
			for j := 0; j < offset; j++ {
				y--
				path = append(path, pos{x, y})
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
