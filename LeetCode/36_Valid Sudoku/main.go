package main

func isValidSudoku(board [][]byte) bool {
	rowCheck := [9][9]bool{}
	columnCheck := [9][9]bool{}
	boxCheck := [9][9]bool{}

	// Номер клуба — то, что стоит на месте во время прохода вдоль клуба.
	// Бегущее - разнообразие внутри клуба.
	for i := 0; i < len(board); i++ { // внутри первой строки (row)
		for j := 0; j < len(board[i]); j++ { // итерируемся по элементам первой строки
			if board[i][j] == '.' { // проверяем точку в байтовом представлении
				continue
			}

			val := board[i][j] - '1' // байты беззнаковые и не имеют отриц. значений

			if rowCheck[i][val] { // rowCheck[клуб][цифра]видели?
				return false
			}

			rowCheck[i][val] = true

			if columnCheck[j][val] {
				return false
			}

			columnCheck[j][val] = true // заполянем в этом же цикле, просто последовательность записи будет нестандатрной - стоя на i строке и читая j элемент на ней, пишем в ПЕРЕВЁРНУТЫЙ массив (грубо говоря в вертикальный, для визуала) значение board[i][j]

			rowOfThree := i / 3
			columnOfThree := j / 3

			checkBox := rowOfThree*3 + columnOfThree

			if boxCheck[checkBox][val] { // всегда array[клуб][цифра]видели?
				return false
			}

			boxCheck[checkBox][val] = true
		}
	}

	return true
}
