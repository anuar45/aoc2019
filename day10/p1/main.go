package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	var m [][]int
	for scanner.Scan() {
		var row []int
		for _, v := range scanner.Text() {
			if v == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		m = append(m, row)
	}

	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m); y++ {
			m[x][y] = m[y][x]
		}
	}
	fmt.Println(m[0][2], m[1][2], m[2][2], m[3][2], m[4][2])
}
