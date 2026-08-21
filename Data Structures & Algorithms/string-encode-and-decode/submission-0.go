type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	for i := 0; i < len(strs); i++ {
		symbolsLen := strconv.Itoa(len(strs[i])) // число -> строка, иначе string(5) - символ по коду

		// identifier
		b.WriteString(symbolsLen) // длина слова
		b.WriteString("#")        // разделитель
		b.WriteString(strs[i])    // сам элемент из входного массива
	}

	return b.String()
}

// 5#Hello3#Bro
func (s *Solution) Decode(encoded string) []string {
	var res []string

	i := 0
	for i < len(encoded) { // пока не дошли до конца (while not in the end, типа). начало гарантированно является identifier-ом, заложенным Encode, поэтому дальше сразу ищем разделитель '#'
		j := i
		for encoded[j] != '#' { // когда стоим на разделителе - выходим. j в этот момент на индексе разделителя
			j++
		}

		symbolsLen, _ := strconv.Atoi(encoded[i:j]) // [i, j) - записанная Encode длина слова после разделителя (не включая правый конец)
		stringElement := encoded[j+1:j+1+symbolsLen]
		res = append(res, stringElement)

		i = j+1+symbolsLen
	}

	return res
}
