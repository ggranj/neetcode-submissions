package main

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, exist := seen[nums[i]]; exist { // если ключ nums[i] есть в мапе - выход
			return true
		}

		seen[nums[i]] = struct{}{} // ключ мапы - значение nums лежащее под индексом i
	}

	return false
}
