package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 4, 7, 70}, 9))
}

// twoSum - O(n) time; O(1) space / memory
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for right > left {
		leftNumber := numbers[left]
		rightNumber := numbers[right]

		sum := rightNumber + leftNumber
		if sum > target {
			right--

			continue
		}

		if sum < target {
			left++

			continue
		}

		return []int{left + 1, right + 1}
	}

	return nil // unreachable by problem guarantee
}
