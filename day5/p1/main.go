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

	var nums []int
	for _, v := range ss {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	for i := 0; i < len(nums); {
		opCode := intToSlice(nums[i])
		fmt.Println(opCode)
		//print(nums)
		if nums[i] == 99 {
			break
		}

		if len(opCode) == 1 {
			switch opCode[len(opCode)-1] {
			case 1:
				nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
				i += 4
			case 2:
				nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
				i += 4
			case 3:
				nums[nums[i+1]] = 1
				i += 2
			case 4:
				fmt.Println(nums[nums[i+1]])
				i += 2
			}
		} else {
			if len(opCode) < 4 {
				reverseInts(opCode)
				opCode = append(opCode, 0)
				reverseInts(opCode)
			}

			var val1, val2, opResult int
			switch opCode[len(opCode)-3] {
			case 0:
				val1 = nums[nums[i+1]]
			case 1:
				val1 = nums[i+1]
			}

			switch opCode[len(opCode)-1] {
			case 3:
				if opCode[len(opCode)-3] == 0 {
					nums[nums[i+1]] = 1
				} else {
					nums[i+1] = 1
				}
				i += 2
				continue
			case 4:
				fmt.Println(val1)
				i += 2
				continue
			}

			switch opCode[len(opCode)-4] {
			case 0:
				val2 = nums[nums[i+2]]
			case 1:
				val2 = nums[i+2]
			}

			switch opCode[len(opCode)-1] {
			case 1:
				opResult = val1 + val2
			case 2:
				opResult = val1 * val2
			}
			// switch opCode[4] {
			// case 0:
			nums[nums[i+3]] = opResult
			i += 4
			// case 1:
			// 	nums[i+3] = opResult
			// }
		}
	}
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

func print(nums []int) {
	for i, v := range nums {
		fmt.Printf(" %d-%d ", i, v)
	}
	fmt.Println()
}
