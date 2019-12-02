package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bb, _ := ioutil.ReadFile("../input.txt")
	ss := strings.Split(string(bb), ",")

	var numsOrig []int
	for _, v := range ss {
		num, _ := strconv.Atoi(v)
		numsOrig = append(numsOrig, num)
	}

	fmt.Println(numsOrig)
	var noun, verb int
loop:
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			nums := make([]int, len(numsOrig))
			copy(nums, numsOrig)
			nums[1] = noun
			nums[2] = verb
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

			if nums[0] == 19690720 {
				fmt.Println(noun, verb)
				break loop
			}
		}
	}
}
