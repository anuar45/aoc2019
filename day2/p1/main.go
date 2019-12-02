package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bb, _ := ioutil.ReadFile("input.txt")
	ss := strings.Split(string(bb), ",")

	var nums []int
	for _, v := range ss {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	nums[1] = 12
	nums[2] = 2

	for i := 0; i < len(nums); i += 4 {
		switch nums[i] {
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		case 99:
			break
		}
	}

	fmt.Println(nums[0])
}
