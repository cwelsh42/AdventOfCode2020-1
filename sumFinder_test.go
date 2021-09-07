package main

import (
	"reflect"
	"testing"
)

func TestFindSumsMultipled(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}
	sumValue := 2020
	expectedAnswer := 514579

	solutionExists, value := findSumsMultiplied(expenseReport, sumValue)

	if !solutionExists {
		t.Errorf("Expected: %t, Received: %t", true, solutionExists)
	}

	if value != expectedAnswer {
		t.Errorf("Expected: %d, Received: %d", expectedAnswer, value)
	}
}

func TestFindSums(t *testing.T) {

	expenseMap := map[int]int{
		1721: 1721,
		979:  979,
		366:  366,
		299:  299,
		675:  675,
		1456: 1456,
	}

	sumValue := 2020

	pairExists, pair := findSums(expenseMap, sumValue)

	if !pairExists {
		t.Errorf("Expected: %t, Received: %t", true, pairExists)
	}

	if (pair[0] != 1721 && pair[1] != 1721) || (pair[0] != 299 && pair[1] != 299) {
		t.Errorf("Expected: %d, %d, Received: %d, %d", 1721, 299, pair[0], pair[1])
	}

}

func TestMakeExpenseMap(t *testing.T) {
	expenseReport := []int{1721, 979, 366, 299, 675, 1456}

	expectedExpenseMap := map[int]int{
		1721: 1721,
		979:  979,
		366:  366,
		299:  299,
		675:  675,
		1456: 1456,
	}

	expenseMap := makeExpenseMap(expenseReport)

	if !reflect.DeepEqual(expenseMap, expectedExpenseMap) {
		t.Errorf("Maps unequal. Expected: %v, Received: %v", expectedExpenseMap, expenseMap)
	}
}
