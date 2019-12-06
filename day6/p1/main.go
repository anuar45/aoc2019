package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("../input.txt")
	data := bufio.NewScanner(f)

	galaxy := make(map[string][]string)

	for data.Scan() {
		orbSet := strings.Split(data.Text(), ")")
		galaxy[orbSet[0]] = append(galaxy[orbSet[0]], orbSet[1])
	}

	var count int
	for planet, orbs := range galaxy {
		for _, _ = range orbs {
			count += CountDepth(galaxy, planet)
		}
	}

	//fmt.Println(galaxy)
	fmt.Println(count)
	//fmt.Println(CountDepth(galaxy, "KWT"))
}

func CountDepth(galaxy map[string][]string, o string) int {
	count := 1
	for planet, orbs := range galaxy {
		for _, orb := range orbs {
			if orb == o {
				count += CountDepth(galaxy, planet)
			}
		}
	}
	return count
}
