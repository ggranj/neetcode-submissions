package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int // пары-тройки создаются по очереди, заранее неизвестно

	sort.Ints(nums) // [-4, -1, -1, 0, 1, 2]

	for i := 0; i <= len(nums)-1; i++ {
		target := nums[i]      // target is an element // target, который надо отрицать. немного ловушка
		left := i + 1          // ниже будем получать как раз элемент через nums[left]
		right := len(nums) - 1 // nums[right]

		for left < right {
			if nums[left]+nums[right] > -target {
				right--

				continue
			}

			if nums[left]+nums[right] < -target {
				left++

				continue
			}

			if nums[left]+nums[right] == -target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				continue
			}

			continue
		}

	}

	return res
}
