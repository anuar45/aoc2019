package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X   int
	Y   int
	Val int
	See []Point
}

type Points struct {
	SizeX  int
	SizeY  int
	Points []Point
}

func main() {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	var points []Point
	var x, y int
	for scanner.Scan() {
		x = 0
		for _, v := range scanner.Text() {
			point := Point{X: x, Y: y}
			if v == '#' {
				point.Val = 1
			}
			points = append(points, point)
			x++
		}
		y++
	}

	pMap := Points{SizeX: x, SizeY: y, Points: points}

	fmt.Println(pMap)
	print(pMap)

	// for y1 := 0; y1 < len(m); y1++ {
	// 	for x1 := 0; x1 < len(m[y1]); x1++ {
	// 		if m[y1][x1] == 1 {
	// 			for y2 := 0; y2 < len(m); y2++ {
	// 				for x2 := 0; x2 < len(m[y2]); x2++ {
	// 					if m[y2][x2] == 1 && y2 != y1 && x2 != x1 {

	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

}

func print(points Points) {
	for i, point := range points.Points {
		fmt.Print(point.Val)
		if i == points.SizeX {
			fmt.Println()
		}
	}
}

func findSlope(x1, y1, x2, y2 int) int {
	var result int
	yDiff := y2 - y1
	xDiff := x2 - x1
	result = yDiff / xDiff
	return result
}
