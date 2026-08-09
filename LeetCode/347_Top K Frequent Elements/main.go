package main

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int, len(nums))  // -10^4 <= nums[i], для отрицательных мапа, тк массив не умеет в отрицательные индексы
	topFrequency := make([][]int, len(nums)+1) // размер = len(nums)+1 (тк индексы - частота появления числа [0] всегда пустой). плюс индексы этого массива являются частотой появления КАКИХ-ТО (может быть несколько чисел) чисел во входном nums. в make не указываем второй аргумент, тк нам нужен слайс заполненный нулями
	result := make([]int, 0, k)

	for i := 0; i < len(nums); i++ {
		frequency[nums[i]]++
	}

	// map[1:3 2:2 3:1]

	for key, val := range frequency { // map[число]сколько раз видели
		topFrequency[val] = append(topFrequency[val], key) // обращение к topFrequency[val] безопасно из-за корректной инициализации
	}

	for i := len(topFrequency) - 1; i >= 0; i-- { // при i = len(topFrequency) сразу же обращение за пределами доступного индекса
		if len(result) == k {
			return result
		}

		for j := range topFrequency[i] {
			result = append(result, topFrequency[i][j])

		}
	}

	return result
}
