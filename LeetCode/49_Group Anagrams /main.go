package main

func groupAnagrams(strs []string) [][]string {
	compareMap := make(map[[26]int][]string, len(strs))

	for i := 0; i < len(strs); i++ {
		passport := [26]int{}

		for j := 0; j < len(strs[i]); j++ { // ходим теперь по конкретному слову внутри strs
			passport[strs[i][j]-'a']++
		}

		// exist не нужен, так как у каждого слова уже свой корректный passport, а нам и не нужно проверять наличие,
		// потому что у каждого слова будет паспорт и оно либо будет в группе с другими, либо будет соло.
		compareMap[passport] = append(compareMap[passport], strs[i])
	}

	result := make([][]string, 0, len(compareMap)) // тут уже знаем нужный cap
	for _, v := range compareMap {
		result = append(result, v)
	}

	return result
}
