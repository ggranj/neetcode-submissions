package main

import "fmt"

func main() {
	fmt.Println('a', 'z', 'A', 'Z', '0', '9')
	fmt.Println('a' - 'A')
	fmt.Println('!')
}

// isPalindrome - O(n) time; O(1) memory / space
func isPalindrome(s string) bool {
	right := len(s) - 1 // for the second pointer
	left := 0

	for right >= left {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

// toLower works with bytes because while reading string like "strs[i]" we're getting bytes.
func toLower(s byte) byte {
	if s >= 'A' && s <= 'Z' { // we need to check input ascii characters, cause isAlphanumeric is more common helper
		return s + ('a' - 'A') // 'a' - 'A' =  32
	}

	return s
}

func isAlphanumeric(elem byte) bool {
	isNumbers := elem >= '0' && elem <= '9'
	isWords := elem >= 'a' && elem <= 'z' || elem >= 'A' && elem <= 'Z'

	return isNumbers || isWords
}
