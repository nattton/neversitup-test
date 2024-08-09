package main

func FindOddNumber(numbers []int) int {
	numberMap := map[int]int{}
	for _, v := range numbers {
		numberMap[v]++
	}

	for k, v := range numberMap {
		if v%2 == 1 {
			return k
		}
	}
	return 0
}
