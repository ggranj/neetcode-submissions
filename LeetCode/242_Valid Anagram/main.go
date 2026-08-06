package main

// string[i] = сырой байт строки из utf-8 // rat == [114, 97, 116]
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s)) // for _, val := range s -> rune (int32, 4 байта)
	tMap := make(map[rune]int, len(t)) // for i := 0;... s[i] -> byte (uint8, 1 байт)

	for _, val := range s {
		sMap[val]++
	}
	for _, val := range t {
		tMap[val]++
	}

	for tKey, tVal := range tMap {
		if sValue, exist := sMap[tKey]; !exist || sValue != tVal {
			return false
		}
	}

	return true
}
