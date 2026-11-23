package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{-2, 0, 0, 2, 2}
	nums3 := []int{0, 0, 0}

	fmt.Println(threeSum(nums1))
	fmt.Println(threeSum(nums2))
	fmt.Println(threeSum(nums3))
}

// threeSum - O(n^2) time, O(1) extra space (in-place sort, output not counted)
func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Ints(nums) // [-4, -1, -1, 0, 1, 2]

	for i := 0; i <= len(nums)-1; i++ {
		target := -nums[i] // the pair on the tail must sum to -nums[i]
		left := i + 1      // we read values as nums[left] / nums[right]
		right := len(nums) - 1

		if i > 0 { // nums[-1] doesn't exist
			if nums[i-1] == nums[i] {
				continue // skip duplicate anchors: the first twin has already done this job; i > 0 guards nums[-1], which doesn't exist
			}
		}

		for left < right { // until the pointers meet (these are indices, not values)
			if nums[left]+nums[right] > target {
				right--

				continue
			}

			if nums[left]+nums[right] < target {
				left++

				continue
			}

			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				left++ // move both: a lone move could only rebuild the same value pair
				right--

				// duplicates are born only at the moment of recording: walk past twins of the just-used values
				for nums[right] == nums[right+1] && left < right {
					right--
				}

				for nums[left] == nums[left-1] && left < right {
					left++
				}

				continue
			}
		}

	}

	return res
}
