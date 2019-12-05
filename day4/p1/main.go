package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")

	params := strings.Split(string(input), "-")

	start, _ := strconv.Atoi(params[0])
	end, _ := strconv.Atoi(params[1])

	var validPasses []int
	for i := start; i <= end; i++ {
		num := intToSlice(i)
		if isAdjacentSame(num) && isIncremSeq(num) {
			validPasses = append(validPasses, i)
		}
	}

	fmt.Println(validPasses)
	fmt.Println(len(validPasses))
}

func isAdjacentSame(nums []int) bool {
	var result bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			result = true
		}
	}
	return result
}

func isIncremSeq(nums []int) bool {
	var result bool
	result = true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			result = false
		}
	}
	return result
}

func intToSlice(n int) []int {
	var intSlice []int
	k := 1
	for n != 0 {
		k *= 10
		d := n % 10
		intSlice = append(intSlice, d)
		n /= 10
	}
	reverseInts(intSlice)
	return intSlice
}

func reverseInts(nums []int) {
	for j, k := 0, len(nums)-1; j < k; {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}
