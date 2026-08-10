package main

func productExceptSelfO1(nums []int) []int { // O(1)
	answer := make([]int, len(nums))

	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		// новый множитель - всегда тот элемент, через который только что перешагнули
		answer[i] = answer[i-1] * nums[i-1]
	}

	multiplier := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * multiplier
		multiplier = multiplier * nums[i]
	}

	return answer
}

func productExceptSelf(nums []int) []int { // O(n)
	left := make([]int, len(nums)) // значение - результат перемножения с предыдущими
	right := make([]int, len(nums))
	answer := make([]int, len(nums))

	left[0] = 1
	right[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1] // left[0] уже задан, используем его и на каждом новом шаге обращаемся к предыдущему результату left[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1] // right[*last index*] тоже задан
	}

	for i := range answer {
		answer[i] = right[i] * left[i]
	}

	return answer
}
