package main

import "fmt"

func main() {
	res := longestConsecutive([]int{100, 1, 2, 3, 0})
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	inputNumsSet := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, exist := inputNumsSet[nums[i]]; !exist {
			inputNumsSet[nums[i]] = true
		}
	}

	finalCounter := 0
	for num := range inputNumsSet {
		if _, exist := inputNumsSet[num-1]; exist { // если есть число меньше нынешнего на 1
			continue
		}
		// здесь мы уже стоим на каком-то num, которое является самым маленьким в контексте непрерывной последовательности n+1, которую мы и будем считать далее

		counter := 1
		for inputNumsSet[num+1] {
			num += 1
			counter++
		}

		if counter > finalCounter {
			finalCounter = counter
		}
	}

	return finalCounter
}
