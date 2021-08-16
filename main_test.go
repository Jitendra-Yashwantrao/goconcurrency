package main

import "testing"

func TestTwoSum_matchfound(t *testing.T) {

	inputNumbers := []int{12, 4, 5, 6, 1, 7}
	output := twoSum(inputNumbers, 7)
	expectedOutPut := []int{3, 4}
	if output[0] != expectedOutPut[0] || output[1] != expectedOutPut[1] {
		t.Errorf("twoSum want %v got %v", expectedOutPut, output)
	}
}
