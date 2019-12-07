package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type node struct {
// 	name    string
// 	friends []*node
// }

func main() {
	f, _ := os.Open("../input.txt")
	data := bufio.NewScanner(f)

	galaxy := make(map[string][]string)

	for data.Scan() {
		orbSet := strings.Split(data.Text(), ")")
		galaxy[orbSet[0]] = append(galaxy[orbSet[0]], orbSet[1])
	}

	var pathA, pathB []string

	detailPath(galaxy, "SAN", &pathA)
	detailPath(galaxy, "YOU", &pathB)

	// find first crossing and distance
	var dist int
Loop:
	for i, v1 := range pathA {
		for j, v2 := range pathB {
			if v1 == v2 {
				dist = i + j
				break Loop
			}
		}
	}

	fmt.Println(dist)

}

func detailPath(galaxy map[string][]string, o string, path *[]string) {
	for planet, orbs := range galaxy {
		for _, orb := range orbs {
			if orb == o {
				*path = append(*path, planet)
				detailPath(galaxy, planet, path)
			}
		}
	}
}
