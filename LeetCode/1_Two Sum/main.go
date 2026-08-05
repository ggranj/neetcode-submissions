package main

func twoSum(nums []int, target int) []int {
	known := make(map[int]int, len(nums)) // [значение]индекс

	for i := 0; i < len(nums); i++ {
		want := target - nums[i]
		if idx, exist := known[want]; exist {
			return []int{i, idx}
		}

		known[nums[i]] = i
	}

	return nil // идиоматичнее, чем result
}
