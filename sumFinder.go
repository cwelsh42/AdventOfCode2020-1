package main

func makeExpenseMap(expenseReport []int) map[int]int {
	expenseMap := make(map[int]int)

	for _, value := range expenseReport {
		expenseMap[value] = value
	}

	return expenseMap
}

func findSums(expenseMap map[int]int, sumValue int) (bool, []int) {
	for _, pair1 := range expenseMap {
		pair2 := sumValue - pair1
		if _, exists := expenseMap[pair2]; exists {
			return true, []int{pair1, pair2}
		}
	}
	return false, []int{}
}

func findSumsMultiplied(expenseReport []int, sumValue int) (bool, int) {
	expenseMap := makeExpenseMap(expenseReport)

	pairExists, pair := findSums(expenseMap, sumValue)

	if pairExists {
		return true, pair[0] * pair[1]
	} else {
		return false, 0
	}
}
