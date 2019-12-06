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
	for k, val1 := range galaxy {
		for i, val2 := range val1 {
			for j, val3 := range galaxy {
				if 
			}
		}
		
		count++
	}

	fmt.Println(orbs)

}
