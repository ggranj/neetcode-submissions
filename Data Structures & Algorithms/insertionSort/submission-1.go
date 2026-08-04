// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	if pairs == nil || len(pairs) == 0 {
		return nil
	}

	localPairs := pairs
	result := make([][]Pair, 0, len(localPairs))
	first := make([]Pair, len(localPairs))
	copy(first, localPairs)
	result = append(result, first)

	for i := 1; i < len(localPairs); i++ { // пропускаем [0], т.к. сравниваем слева-направо
		cur := localPairs[i] // взяли [1] - banana
		prev := i - 1        // индекс предыдущей от cur пары для сравнения
		
		for prev >= 0 && localPairs[prev].Key > cur.Key { // до тех пор, пока prev >= 0 && localPairs[prev].Key > cur.Key
			localPairs[prev+1] = localPairs[prev] // выполняется это
			prev--                                // и это
		}

		localPairs[prev+1] = cur
		snap := make([]Pair, len(localPairs))
		copy(snap, localPairs)
		result = append(result, snap)
	}

	return result
}
