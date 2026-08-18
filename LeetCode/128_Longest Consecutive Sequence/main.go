package main

func longestConsecutive(nums []int) int {
	inputNumsSet := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, exist := inputNumsSet[nums[i]]; !exist {
			inputNumsSet[nums[i]] = true
		}
	}

	finalCounter := 0
	for num := range inputNumsSet {
	}

	return finalCounter
}
