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
	for _, v := range galaxy {
		for _, o := range v {
			count += CountDepth(galaxy, o)
		}
	}

	fmt.Println(CountDepth(galaxy, "W4F"))
	fmt.Println(count)

}

func CountDepth(galaxy map[string][]string, orb string) int {
	var count int
	for k, v := range galaxy {
		for _, o := range v {
			if o == orb {
				count += CountDepth(galaxy, k)
			}
		}
	}
	return count
}
